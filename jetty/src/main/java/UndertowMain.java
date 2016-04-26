import io.undertow.servlet.api.DeploymentInfo;
import io.undertow.servlet.api.DeploymentManager;
import io.undertow.servlet.Servlets;
import io.undertow.server.handlers.PathHandler;
import io.undertow.Handlers;
import io.undertow.Undertow;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class UndertowMain {
  /**
   * Main.main is the application entry point.
   *
   * @param args command line arguments
   * @throws java.lang.Exception Blow up when some exception is raised here.
   */
  public static void main(String[] args) throws Exception {
    DeploymentInfo servletBuilder = Servlets.deployment()
        .setClassLoader(UndertowMain.class.getClassLoader())
        .setContextPath("/")
        .setDeploymentName("test.war")
        .addServlets(
            Servlets.servlet("HelloServlet", HelloServlet.class).addMapping("/*")
        );
    
    DeploymentManager manager = Servlets.defaultContainer().addDeployment(servletBuilder);
    manager.deploy();

    PathHandler path = Handlers.path(Handlers.redirect("/"))
        .addPrefixPath("/", manager.start());

    Undertow server = Undertow.builder()
        .addHttpListener(8080, "localhost")
        .setHandler(path)
        .build();
    server.start();
  }

  public static class HelloServlet extends HttpServlet {
    @Override protected void doGet(HttpServletRequest request, HttpServletResponse response)
      throws ServletException, IOException {
      response.setContentType("text/plain");
      response.setStatus(HttpServletResponse.SC_OK);
      response.getWriter().println("Hello World from servlet");
    }
  }
}
