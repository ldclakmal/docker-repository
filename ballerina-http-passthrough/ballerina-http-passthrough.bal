import ballerina/http;
import ballerina/log;

listener http:Listener passthroughListener = new(9090);

http:Client passthroughClient = new("http://localhost:9191");

@http:ServiceConfig { basePath: "/passthrough" }
service passthroughService on passthroughListener {

    @http:ResourceConfig { path: "/" }
    resource function passthrough(http:Caller outboundEP, http:Request clientRequest) {
        var response = passthroughClient->forward("/hello/sayHello", clientRequest);
        if (response is http:Response) {
            checkpanic outboundEP->respond(response);
        } else {
            log:printError("Error at passthrough service", err = response);
            http:Response res = new;
            res.statusCode = http:STATUS_INTERNAL_SERVER_ERROR;
            json errMsg = { message: <string>response.detail()?.message };
            res.setPayload(errMsg);
            checkpanic outboundEP->respond(res);
        }
    }
}
