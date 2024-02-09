
## [![Mentioned in Awesome OpenVINO](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/openvinotoolkit/awesome-openvino)
## The flask app
The Flask code you've provided is for a web application that predicts molecular properties based on SMILES (Simplified Molecular Input Line Entry System) strings. Here's a breakdown of what each part of the code does:

Importing Libraries: The code imports necessary libraries, including Flask for web application functionalities, RDKit for cheminformatics tasks, NumPy for numerical operations, and OpenVINO for running inference with a pre-trained model.

Initializing Flask App: A Flask application instance is created.

Loading OpenVINO Model: The application loads a pre-trained OpenVINO model from the specified path ('lipophilicity_openvino.xml'). This model is used to predict properties of molecules.

Function smiles_to_fp:

This function converts a SMILES string to a molecular fingerprint, which is a vector representation of the molecule.
The fingerprint is generated using RDKit's GetMorganFingerprintAsBitVect method, specifying a radius of 2 and a bit vector size (n_bits) of 2048.
The fingerprint is returned as a NumPy array, suitable for use as input to the OpenVINO model.
API Endpoint /predict:

The Flask app defines an endpoint (/predict) that handles POST requests.
When a request is received, it extracts the SMILES string from the JSON data in the request body.
The SMILES string is converted to a molecular fingerprint using the smiles_to_fp function.
This fingerprint is then fed into the OpenVINO model for inference.
The result of the model prediction is returned as a JSON response.
Running the App:

The Flask application is configured to run in debug mode on port 5000.
In summary, this web application provides an API endpoint for predicting molecular properties using a machine learning model. A client can send a SMILES string to this endpoint, and the server processes this input to generate a prediction, which is then sent back to the client. The use of RDKit for chemical informatics and OpenVINO for efficient model inference makes this application suitable for tasks in computational chemistry and drug discovery.

## Go Client
command-line client application designed to interact with the Flask API for predicting molecular properties based on SMILES (Simplified Molecular Input Line Entry System) strings. Here's a breakdown of its functionality:

Importing Packages: It imports necessary Go packages for handling HTTP requests, JSON processing, standard input/output operations, and file handling.

PredictionRequest Structure:

A struct PredictionRequest is defined with a field Smiles, which is annotated with a JSON struct tag. This struct is used to format the JSON payload for the API request.
Function makePredictionRequest:

This function takes the API URL and a SMILES string as input.
It creates a PredictionRequest instance with the provided SMILES string and marshals it into JSON.
A POST request is made to the specified API URL with this JSON data.
It assumes that the server's response is an image file and attempts to save this file locally as molecule_image.png.
The function includes error handling for JSON marshaling, HTTP request errors, file creation, and file writing.
Main Function:

The main function defines the Flask API URL (http://localhost:5000/predict).
It prompts the user to enter a SMILES string via the command line.
Once the user inputs a SMILES string, the function makePredictionRequest is called with this string and the API URL.
The main purpose is to send the entered SMILES string to the Flask API and save the response (expected to be an image) locally.
In summary, this Go application serves as a frontend interface for the Flask backend. It allows a user to input a SMILES string and communicates with the Flask API to get a prediction. The prediction results (in this case, an image representing the molecule) are saved locally by the Go application, offering a simple and interactive way for users to utilize the backend's capabilities.

## Complete Workflow
The complete workflow of the combined Flask (Python) and Go applications creates a system for predicting molecular properties and retrieving molecule images using SMILES strings. Here's how the entire process flows:

Flask Backend:
Initialization: The Flask application starts and initializes the OpenVINO model for molecular property prediction.

Endpoint Setup: The /predict endpoint is set up to handle POST requests containing SMILES strings in JSON format.

Request Handling:

When a request is received, the Flask app extracts the SMILES string from the JSON payload.
It converts the SMILES string to a molecular fingerprint using RDKit.
This fingerprint is fed into the OpenVINO model to predict the molecular properties.
Response: The Flask app sends back the prediction results, typically numerical properties of the molecule derived from the model.

Go Frontend:
User Interaction: The Go application runs in the command line, prompting the user to enter a SMILES string.

Sending Request:

The Go app takes the user input and creates a JSON payload with the SMILES string.
It sends this payload to the Flask API's /predict endpoint via a POST request.
Receiving and Handling Response:

The Go app receives the response from the Flask API.
Assuming the response is an image file (as per your initial assumption, although the Flask code provided doesn't seem to return an image), the Go application saves this file locally.
Complete Workflow:
Start: The user starts the Go application and enters a SMILES string when prompted.

Processing Request: The Go application sends this SMILES string to the Flask backend.

Backend Processing:

The Flask app receives the SMILES string, processes it to generate a molecular fingerprint, and uses a machine learning model to predict properties.
The Flask app sends the prediction results back to the Go application.
Receiving Output: The Go application receives the prediction results and potentially saves any returned data (like an image) locally.

End Result: The user is provided with the molecular property prediction and any additional output (like a molecule image) saved on their local system.

This system leverages the strengths of both Flask and Go: Flask handles complex backend processing involving cheminformatics and machine learning, while the Go application offers a user-friendly interface to easily input data and receive results.
