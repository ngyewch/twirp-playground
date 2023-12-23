import io.helidon.webserver.Routing;
import io.helidon.webserver.WebServer;
import rpc.AddRequest;
import rpc.AddResponse;
import rpc.RpcTwirp;
import rpc2.ToUpperRequest;
import rpc2.ToUpperResponse;

public class Main {
  public static void main(String[] args) throws Exception {
    final WebServer webServer =
        WebServer.builder()
            .port(8080)
            .routing(
                Routing.builder()
                    .register(
                        "/twirp",
                        rules -> {
                          RpcTwirp.Helidon.TestService.update(rules, new TestService());
                          RpcTwirp.Helidon.TestService2.update(rules, new TestService2());
                        })
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

  public static class TestService2 implements RpcTwirp.TestService2 {

    @Override
    public ToUpperResponse toUpper(ToUpperRequest input) {
      return ToUpperResponse.newBuilder().setText(input.getText().toUpperCase()).build();
    }
  }
}
