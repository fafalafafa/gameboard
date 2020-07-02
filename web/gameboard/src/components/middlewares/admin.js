import Axios from "axios";
import { GetHost } from "../libs/gethost";
import qs from "querystring"

export const GetCharacters = () => {
    return Axios.get(`http://${GetHost()}/admin/get-all-characters`)
}

export const StartNewGame = () => {
    return Axios.get(`http://${GetHost()}/admin/start-game`)
}
