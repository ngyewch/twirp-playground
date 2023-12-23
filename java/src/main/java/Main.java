import io.helidon.webserver.Routing;
import io.helidon.webserver.WebServer;
import rpc.AddRequest;
import rpc.AddResponse;
import rpc.RpcTwirp;

public class Main {
  public static void main(String[] args) throws Exception {
    final WebServer webServer =
        WebServer.builder()
            .port(8080)
            .routing(
                Routing.builder()
                    .register(
                        "/twirp",
                        rules -> RpcTwirp.Helidon.TestService.update(rules, new TestService()))
                    .build())
            .build();
    webServer.start();
  }

  public static class TestService implements RpcTwirp.TestService {

    @Override
    public AddResponse add(AddRequest input) {
      return AddResponse.newBuilder().setValue(input.getA() + input.getB()).build();
    }
  }
}
