import pandas as pd
import psycopg2
from dotenv import load_dotenv
import os
import logging

# Configurar o logger
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

# Carregar as variáveis de ambiente do arquivo .env
load_dotenv()

# Obter as variáveis de ambiente
DB_HOST = os.getenv('DB_HOST')
DB_PORT = os.getenv('DB_PORT')
DB_USER = os.getenv('DB_USER')
DB_PASSWORD = os.getenv('DB_PASSWORD')
DB_NAME = os.getenv('DB_NAME')
DB_SSLMODE = os.getenv('DB_SSLMODE')

# Carregar o DataFrame
df_ratings = pd.read_csv('ratings.csv')
logger.info('DataFrame carregado com sucesso')

# Conectar ao banco de dados PostgreSQL para obter os user_ids válidos
def obter_user_ids_validos():
    logger.info('Conectando ao banco de dados para obter os user_ids válidos')
    conn = psycopg2.connect(
        host=DB_HOST,
        port=DB_PORT,
        user=DB_USER,
        password=DB_PASSWORD,
        dbname=DB_NAME,
        sslmode=DB_SSLMODE
    )
    cur = conn.cursor()
    cur.execute("SELECT id FROM users")
    user_ids = [row[0] for row in cur.fetchall()]
    cur.close()
    conn.close()
    logger.info(f'{len(user_ids)} user_ids válidos obtidos')
    return user_ids

# Obter os user_ids válidos
user_ids_validos = obter_user_ids_validos()

# Filtrar para remover ratings com project_id maior que 1410 e user_id não presentes em user_ids_validos
logger.info('Filtrando DataFrame')
df_ratings = df_ratings[(df_ratings['project_id'] <= 1410) & (df_ratings['user_id'].isin(user_ids_validos))]
logger.info(f'DataFrame filtrado para {len(df_ratings)} entradas válidas')

# Convertendo colunas para tipos nativos do Python
df_ratings['user_id'] = df_ratings['user_id'].astype(int)
df_ratings['project_id'] = df_ratings['project_id'].astype(int)
df_ratings['rating'] = df_ratings['rating'].astype(int)

# Função para conectar ao banco de dados e inserir os dados
def inserir_dados(df):
    logger.info('Inserindo dados no banco de dados')
    # Conectar ao banco de dados PostgreSQL
    conn = psycopg2.connect(
        host=DB_HOST,
        port=DB_PORT,
        user=DB_USER,
        password=DB_PASSWORD,
        dbname=DB_NAME,
        sslmode=DB_SSLMODE
    )

    # Criar um cursor
    cur = conn.cursor()

    # Inserir os dados na tabela ratings
    insert_query = "INSERT INTO ratings (user_id, project_id, rating) VALUES (%s, %s, %s)"
    for index, row in df.iterrows():
        try:
            cur.execute(insert_query, (int(row['user_id']), int(row['project_id']), int(row['rating'])))
        except Exception as e:
            logger.error(f'Erro ao inserir dados na linha {index}: {e}')
    
    # Confirmar as transações
    conn.commit()

    # Fechar o cursor e a conexão
    cur.close()
    conn.close()
    logger.info('Dados inseridos com sucesso')

# Inserir os dados do DataFrame no banco de dados
inserir_dados(df_ratings)

