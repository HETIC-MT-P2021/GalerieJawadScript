module Route exposing (Route (..), matchRoute)

import Url exposing (Url)
import Url.Parser exposing (..)

type Route
    = NotFound
    | Home
    | Image Int
    | AddImage
    | EditImage
    | ImageByTag String
    | ImageByCategory String
    | Categories
    | Tags

matchRoute : Parser (Route -> a) a
matchRoute =
    oneOf
        [ map Home top
        , map Image ( s "image" </> int )
        , map AddImage ( s "new" )
        , map EditImage ( s "edit" )
        , map ImageByTag ( s "tags" </> string )
        , map ImageByCategory ( s "category" </> string )
        , map Categories (s "categories" )
        , map Tags (s "tags" )
        ]


parseUrl : Url -> Route
parseUrl url =
    case parse matchRoute url of
        Just route ->
            route

        Nothing ->
            NotFound