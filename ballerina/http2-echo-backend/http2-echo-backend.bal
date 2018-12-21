import ballerina/http;
import ballerina/log;
import ballerina/mime;

listener http:Listener listenerEP = new(9191, config = { httpVersion: "2.0" });

service hello on listenerEP {

    resource function sayHello(http:Caller outboundEP, http:Request clientRequest) {
        http:Response res = new;
        var payload = clientRequest.getJsonPayload();
        if (payload is json) {
            res.setPayload(untaint payload);
        } else {
            string errMsg = "Received payload is not json compatible";
            log:printError(errMsg, err = payload);
            res.setPayload({ message: errMsg });
        }
        res.setHeader(http:CONTENT_TYPE, mime:APPLICATION_JSON);
        _ = outboundEP->respond(res);
    }
}
