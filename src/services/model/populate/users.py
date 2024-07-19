import pandas as pd
import json
import subprocess

# Carregar o CSV em um DataFrame
df = pd.read_csv('users.csv')

# Função para criar o payload JSON
def criar_payload(linha):
    payload = {
        "name": linha['proponent'],
        "email": linha['email'],
        "password": linha['password'],
        "company_name": linha['company name'],
        "office": linha['position'],
        "linkedin_link": linha['linkedin_link'],
        "interest": linha['interest']
    }
    return json.dumps(payload)

# Criar payloads para cada linha e executar o curl
for index, linha in df.iterrows():
    payload = criar_payload(linha)
    # Montar o comando curl
    command = f"curl -X POST -H 'Content-Type: application/json' -d '{payload}' http://localhost:8082/api/v1/register"
    # Executar o comando usando subprocess
    subprocess.run(command, shell=True)

