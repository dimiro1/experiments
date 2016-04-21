import java.io.InputStream;
import java.io.FileInputStream;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;

public class Main {
    public static void main(String[] args) {
        try {
            InputStream in = new FileInputStream("Main.java");
            BufferedReader reader = new BufferedReader(new InputStreamReader(in));
            StringBuilder builder = new StringBuilder();
            String line;
            
            while ((line = reader.readLine()) != null) {
                builder.append(line).append("\n");
            }
            
            reader.close();
            System.out.println(builder.toString());
        } catch (IOException e) {}
    }
}