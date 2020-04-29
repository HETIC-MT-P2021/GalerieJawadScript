module Main exposing (..)

import Browser
import File exposing (File)
import File.Select as Select
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Attributes exposing (src)
import Html.Events exposing (..)
import Json.Decode as D
import Task

import File
import Http
import Json.Decode exposing (list, string)

import Header
import Footer



-- MAIN

foof = 
  Footer.footerSauce

heeh = 
  Header.headerSauce


main =
  Browser.element
    { init = init
      , view = view
      , update = update
      , subscriptions = subscriptions
    }



-- MODEL


type alias Model =
  { hover : Bool
  , previews : List String
  }


init : () -> (Model, Cmd Msg)
init _ =
  (Model False [], Cmd.none)



-- UPDATE


type Msg
  = Pick
  | DragEnter
  | DragLeave
  | GotFiles File (List File)
  | GotPreviews (List String)
  | Modify
  -- | Post

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    Pick ->
      ( model
      , Select.files ["image/*"] GotFiles
      )

    DragEnter ->
      ( { model | hover = True }
      , Cmd.none
      )

    DragLeave ->
      ( { model | hover = False }
      , Cmd.none
      )

    GotFiles file files ->
      ( { model | hover = False }
      , Task.perform GotPreviews <| Task.sequence <|
          List.map File.toUrl (file :: files)
      )

    GotPreviews urls ->
      ( { model |  previews = model.previews ++ urls }
      , Cmd.none
      )
    
    Modify ->
      ( model
      , Select.files ["image/*"] GotFiles
      )

-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  div
  [ style "background" "#010a11"
  , style "width" "100%"
  , style "height" "100%"
  , style "overflow" "hidden"
  ]
  [
    heeh
    ,div
      [ style "width" "80%"
      , style "margin" "100px auto"
      , style "padding" "40px"
      , style "display" "flex"
      , style "flex-direction" "column"
      , style "justify-content" "center"
      , style "align-items" "center"
      , hijackOn "dragenter" (D.succeed DragEnter)
      , hijackOn "dragover" (D.succeed DragEnter)
      , hijackOn "dragleave" (D.succeed DragLeave)
      , hijackOn "drop" dropDecoder
      ]
      [ button [ onClick Pick ] [ text "Upload New Image" ]
      , div
          [ style "padding" "40px"
          , style "display" "flex"
          , style "justify-content" "center"
          , style "align-items" "center"
          , style "flex-wrap" "wrap"
          ]
          (List.map imagePreview model.previews)
      ]
    , foof
  ]


imagePreview : String -> Html Msg
imagePreview url =
  div[]
  [
  div
    [ class "category" 
    , style "width" "100px"
    , style "height" "100px"
    , style "background-image" ("url('" ++ url ++ "')")
    , style "background-position" "center"
    , style "background-repeat" "no-repeat"
    , style "background-size" "contain"
    , style "display" "flex"
    , style "flex-direction" "column"
    , style "justify-content" "center"
    , style "align-items" "center"
    , style "flex-wrap" "wrap"
    , style "margin" "20px"
    ]
    [
      button [ onClick Modify 
      , style "margin-top" "110px"
      ] 
      [ text "Modify Image" ]
    ]
  ]


dropDecoder : D.Decoder Msg
dropDecoder =
  D.at ["dataTransfer","files"] (D.oneOrMore GotFiles File.decoder)


hijackOn : String -> D.Decoder msg -> Attribute msg
hijackOn event decoder =
  preventDefaultOn event (D.map hijack decoder)


hijack : msg -> (msg, Bool)
hijack msg =
  (msg, True)

