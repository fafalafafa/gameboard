export const GetHost = () => {
    let { hostname, host } = window.location;
    if ( host.indexOf(":3000") >=0 ) {
        return `${hostname}:8080`
    }
    return hostname;
}