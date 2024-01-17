from flask import Flask, request, jsonify
from rdkit import Chem
from rdkit.Chem import AllChem
import numpy as np
import openvino.runtime as ov

app = Flask(__name__)

# Load OpenVINO model
model_path = 'lipophilicity_openvino.xml'  # Replace with your model path
core = ov.Core()
compiled_model = core.compile_model(model_path, "CPU")

# Function to convert SMILES to fingerprints
def smiles_to_fp(smiles, n_bits=2048):
    mol = Chem.MolFromSmiles(smiles)
    fp = AllChem.GetMorganFingerprintAsBitVect(mol, radius=2, nBits=n_bits)
    return np.array(fp)

# Endpoint for predictions
@app.route('/predict', methods=['POST'])
def predict():
    data = request.json
    smiles = data['smiles']
    fp = smiles_to_fp(smiles)
    input_tensor = np.array([fp], dtype=np.float32)

    # OpenVINO inference
    ov_input_tensor = ov.Tensor(input_tensor)
    result = compiled_model([ov_input_tensor])[0]
    prediction = result[0]

    return jsonify({'prediction': prediction.tolist()})

if __name__ == '__main__':
    app.run(debug=True, port=5000)
