import ballerina/http;
import ballerina/log;
import ballerina/mime;

listener http:Listener listenerEP = new(9191, config = { httpVersion: "2.0" });

service hello on listenerEP {

    resource function sayHello(http:Caller outboundEP, http:Request clientRequest) {
        http:Response res = new;
        var entity = clientRequest.getEntity();
        if (entity is mime:Entity) {
            res.setEntity(entity);
        } else {
            string errMsg = "An error occurred while retrieving the entity from the backend";
            log:printError(errMsg, err = entity);
            res.setPayload({ message: errMsg });
        }
        res.setHeader(http:CONTENT_TYPE, mime:APPLICATION_JSON);
        checkpanic outboundEP->respond(res);
    }
}
