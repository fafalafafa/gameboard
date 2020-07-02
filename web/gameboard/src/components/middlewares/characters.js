import Axios from "axios";
import { GetHost } from "../libs/gethost";
import qs from "querystring"

export const GetCharacters = () => {
    return Axios.get(`http://${GetHost()}/admin/get-all-characters`)
}

export const GetCharacter = (name) => {
    return Axios.post(`http://${GetHost()}/get-character`, qs.stringify({
        name: name
    }));
}

export const GetCharacterDetail = (sessionId) => {
    let _qs = qs.stringify({
        sessionId: sessionId
    });
    return Axios.get(`http://${GetHost()}/get-character?${_qs}`);
}