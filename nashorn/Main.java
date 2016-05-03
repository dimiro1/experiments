import java.io.FileReader;
import java.io.FileNotFoundException;

import javax.script.Invocable;
import javax.script.ScriptEngine;
import javax.script.ScriptEngineManager;
import javax.script.ScriptException;

public class Main {
    /**
     * The main application entry point
     *
     * @param args application arguments
     * @throws ScriptException in case of an error cased by a script
     * @throws FileNotFoundException in case the config file could not be found
     */
    public static void main(String[] args) throws ScriptException, FileNotFoundException {
        ScriptEngine engine = new ScriptEngineManager().getEngineByName("nashorn");
        engine.eval(new FileReader("config.js"));
        
        System.out.println(engine.get("endpoint"));
        System.out.println(engine.get("protocol"));
        
        Object obj = engine.get("helloService");
        
        Invocable invocable = (Invocable) engine;
        HelloService helloService = invocable.getInterface(obj, HelloService.class);
        
        System.out.println(helloService.hello("Claudemiro"));
    }
}