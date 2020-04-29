module Footer exposing (..)

import Browser
import File exposing (File)
import File.Select as Select
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Attributes exposing (src)
import Html.Events exposing (..)
import Json.Decode as D


init = 
    { sauce = 0 }

footerSauce : Html msg
footerSauce =
    div 
    [ style "width" "100%"
    , style "background" "#00121b"
    , style "display" "block"   
    ] 
    [
        div
        [ style "width" "95%"
        , style "margin" "auto"
        , style "padding" "30px 10px"
        , style "display" "flex"
        , style "flex-wrap" "wrap"
        , style "box-sizing" "border-box"
        , style "justify-content" "center"
        ]
        [
           div
           [ style "width" "25%"
           , style "padding" "10px 20px"
           , style "box-sizing" "border-box"
           , style "color" "#fff" 
           ] 
           [
               h1
               [ style "color" "#fff"]
               [text "Galerie Jawadscript"]
           
                ,p
                [ style "font-size" "16px"
                , style "text-align" "justify"
                , style "line-height" "25px"
                , style "color" "#fff"
                ]
                [ text "An image Galery with a very bad sense of humor"]
           ]
           ,div
           [ style "width" "25%"
           , style "padding" "10px 20px"
           , style "box-sizing" "border-box"
           , style "color" "#fff" 
           ] 
           [

                div
                [ style "color" "#fff"]
                [
                    h3
                    []
                    [ text "Quick Links"]
                    ,div
                    [ style "height" "3px"
                    , style "width" "40px"
                    , style "background" "#ff9800"
                    , style "color" "#ff9800"
                    , style "background-color" "#ff9800"
                    , style "border" "0px"
                    ]
                    []
                    ,ul
                    [ style "list-style" "none"
                    , style "color" "#fff"
                    , style "font-size" "15px"
                    , style "letter-spacing" "0.5px"
                    ]
                    [
                        a
                        [ style "text-decoration" "none"
                        , style "outline" "none"
                        , style "color" "#fff"
                        , style "transition" "0.3s"
                        ]
                        [
                            li
                            [ style "margin" "10px 0"
                            , style "height" "25px"
                            ]
                            [text "Home"]
                        ]
                        ,a
                        [ style "text-decoration" "none"
                        , style "outline" "none"
                        , style "color" "#fff"
                        , style "transition" "0.3s"
                        ]
                        [
                            li
                            [ style "margin" "10px 0"
                            , style "height" "25px"
                            ]
                            [text "Categories"]
                        ]
                        ,a
                        [ style "text-decoration" "none"
                        , style "outline" "none"
                        , style "color" "#fff"
                        , style "transition" "0.3s"
                        ]
                        [
                            li
                            [ style "margin" "10px 0"
                            , style "height" "25px"
                            ]
                            [text "Galery"]
                        ]
                        ,a
                        [ style "text-decoration" "none"
                        , style "outline" "none"
                        , style "color" "#fff"
                        , style "transition" "0.3s"
                        ]
                        [
                            li
                            [ style "margin" "10px 0"
                            , style "height" "25px"
                            ]
                            [text "Search"]
                        ]
                    ]
                ]
           ]
        ]
    ]


view model = 
    footerSauce
    

update model = 
    model

main = 
    Browser.sandbox
    {
        init = init
        , view = view
        , update = update
    }

