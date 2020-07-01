import axios from "axios";

axios.interceptors.request.use(req => {
    // `req` is the Axios request config, so you can modify
    // the `headers`.
    req.headers["Content-Type"]= "application/x-www-form-urlencoded;charset=utf-8";
    return req;
});
  