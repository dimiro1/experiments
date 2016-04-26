import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.servlet.ServletHandler;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class Main {
  /**
   * Main.main is the application entry point.
   *
   * @param args command line arguments
   * @throws java.lang.Exception Blow up when some exception is raised here.
   */
  public static void main(String[] args) throws Exception {
    Server server = new Server(8080);

    ServletHandler handler = new ServletHandler();
    server.setHandler(handler);

    handler.addServletWithMapping(HelloServlet.class, "/*");

    server.start();
    server.join();
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
