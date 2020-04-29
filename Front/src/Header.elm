module Header exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Attributes exposing (src)
import Html.Events exposing (..)


init = 
    { sauce = 0 }

headerSauce : Html msg
headerSauce =
    div 
    [ style "box-sizing" "border-box"]
    [
        div
        [ style "height" "50px"
        , style "width" "100%"
        , style "background-color" "#00121b"
        , style "position" "relative"
        ]
        [
            div
            [ style "display" "inline"]
            [
                div
                [ style "display" "inline-block"
                , style "font-size" "22px"
                , style "color" "#fff"
                , style "padding" "10px 10px 10px 10px"
                ]
                [text "Galery Jawadscript"]
            ]

            , div
            [ style "display" "inline"
            , style "float" "right"
            , style "font-size" "18px"
            ]
            [
                a
                [ style "display" "inline-block"
                , style "padding" "10px 10px 13px 10px"
                , style "text-decoration" "none"
                , style "color" "#efefef"]
                [ text "Github"]
                ,a
                [ style "display" "inline-block"
                , style "padding" "10px 10px 13px 10px"
                , style "text-decoration" "none"
                , style "color" "#efefef"]
                [ text "Categories"]
                ,a
                [ style "display" "inline-block"
                , style "padding" "10px 10px 13px 10px"
                , style "text-decoration" "none"
                , style "color" "#efefef"]
                [ text "Search"]
            ]
        ]
        
    ]


view model = 
    headerSauce
    

update model = 
    model

main = 
    Browser.sandbox
    {
        init = init
        , view = view
        , update = update
    }

