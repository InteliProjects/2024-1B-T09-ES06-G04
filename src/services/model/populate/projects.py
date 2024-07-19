import pandas as pd
import json
import subprocess

# Carregar o DataFrame df_projects (substitua pelo caminho correto do seu arquivo)
df_projects = pd.read_csv('projects.csv')

# Função para criar o payload JSON
def criar_payload(linha):
    payload = {
        "name": linha['project'],  # Nome do projeto
        "description": "Descrição do projeto",  # Descrição fictícia
        "macro_setor": linha['macrosector'],  # Macro setor do projeto
        "micro_setor": "Setor específico",  # Setor específico fictício
        "image_link": "http://example.com/image.jpg",  # Link de imagem fictício
        "user_id": linha['proponent_id']  # ID do proponente como user_id
    }
    return json.dumps(payload)

# Criar payloads para cada linha e executar o curl
for index, linha in df_projects.iterrows():
    payload = criar_payload(linha)
    # Montar o comando curl
    command = f"curl -X POST http://localhost:8082/api/v1/projects -H 'Content-Type: application/json' -d '{payload}'"
    # Executar o comando usando subprocess
    subprocess.run(command, shell=True)

