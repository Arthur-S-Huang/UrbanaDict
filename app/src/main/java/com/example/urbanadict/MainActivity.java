package com.example.urbanadict;

import androidx.appcompat.app.AppCompatActivity;
import android.os.Bundle;
import android.widget.Button;
import android.widget.TextView;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
//import java.util.*;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Button search = findViewById(R.id.search);
        search.setOnClickListener(unused -> searchDictionary());
    }

    public void searchDictionary() {
        TextView emailEntered = findViewById(R.id.editText);
        TextView content = findViewById(R.id.content);
        String input = emailEntered.getText().toString();
        try {
            Process process = Runtime.getRuntime().exec("go run urban.go " + input);

            StringBuilder output = new StringBuilder();

            BufferedReader reader = new BufferedReader(
                    new InputStreamReader(process.getInputStream()));

            String line;
            while ((line = reader.readLine()) != null) {
                output.append(line + "\n");
            }

            int exitVal = process.waitFor();
            if (exitVal == 0) {
                content.setText(output);
            }
        } catch (Exception e) {
            content.setText("error searching for word");
        }
    }
}

