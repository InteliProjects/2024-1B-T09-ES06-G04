from flask import Flask, jsonify, request
import psycopg2
import base64
import json
from dotenv import load_dotenv
import os
import numpy as np
from sklearn.neighbors import KNeighborsClassifier
from sklearn.preprocessing import LabelEncoder
from random import shuffle

# Load environment variables from .env file
load_dotenv()

app = Flask(__name__)

# Get database connection details from environment variables
DB_HOST = os.getenv('DB_HOST')
DB_NAME = os.getenv('DB_NAME')
DB_USER = os.getenv('DB_USER')
DB_PASSWORD = os.getenv('DB_PASSWORD')

# Function to connect to the database
def connect_to_db():
    conn = psycopg2.connect(
        host=DB_HOST,
        database=DB_NAME,
        user=DB_USER,
        password=DB_PASSWORD
    )
    return conn

# Function to get user ID from request header
def get_id_from_header_request(request):
    token_auth = request.headers.get("Authorization")
    if token_auth:
        token = token_auth.replace("Bearer ", "")
        try:
            # Decode the JWT token without verifying the signature
            parts = token.split('.')
            if len(parts) == 3:
                claims_data = base64.urlsafe_b64decode(parts[1] + '==')
                claims = json.loads(claims_data)
                if 'id' in claims:
                    return claims['id']
        except Exception as e:
            return ""

    token_query = request.args.get("token")
    if token_query:
        return token_query

    return ""

# Function to calculate Jaccard similarity
def jaccard_similarity(set1, set2):
    intersection = len(set1.intersection(set2))
    union = len(set1.union(set2))
    if union == 0:
        return 0.0
    return intersection / union

# Function to calculate similarity based on the same interest
def calculate_similarity_same_interest(user_id, users_with_interest, user_ratings):
    similarity_scores = {}

    # Check if the user ID is present in the user_ratings dictionary
    if user_id in user_ratings:
        user_projects = set(user_ratings[user_id].keys())

        for other_user_id, other_user_interest in users_with_interest:
            if other_user_id != user_id:
                # Check if the other user's ID is present in the user_ratings dictionary
                if other_user_id in user_ratings:
                    other_user_projects = set(user_ratings[other_user_id].keys())
                    similarity = jaccard_similarity(user_projects, other_user_projects)
                    similarity_scores[other_user_id] = similarity
                else:
                    # If the other user's ID is not present, set similarity to 0
                    similarity_scores[other_user_id] = 0

        sorted_similarity_scores = sorted(similarity_scores.items(), key=lambda x: x[1], reverse=True)
        return sorted_similarity_scores
    else:
        # If the user ID is not present, return an empty list of similarities
        return []

# Function to recommend all projects based on probability
def recommend_all_projects_prob(user_id, user_ratings, k=3):
    # Check if there are user ratings
    if not user_ratings.get(user_id):
        return []

    # Connect to the database
    conn = connect_to_db()
    cursor = conn.cursor()

    # Calculate similarity with other users having the same interest
    query = "SELECT id, interest FROM users WHERE id = %s;"
    cursor.execute(query, (user_id,))
    user_interest = cursor.fetchone()
    if user_interest:
        user_interest_name = user_interest[1]

        query = "SELECT id, interest FROM users WHERE interest = %s AND id != %s;"
        cursor.execute(query, (user_interest_name, user_id))
        users_with_same_interest = cursor.fetchall()

        # Calculate similarity between the specified user and other users with the same interest
        user_similarity_scores = calculate_similarity_same_interest(user_id, users_with_same_interest, user_ratings)

        # Get all available projects
        all_projects = list(user_ratings.get(user_id).keys())

        # Create a list of tuples with project data and their ratings by similar users
        df_similar_projects = [(other_user_id, project_id, rating)
                               for other_user_id, _ in user_similarity_scores
                               for project_id, rating in user_ratings.get(other_user_id, {}).items()]

        if not df_similar_projects:
            # Close the database connection
            conn.close()
            return []

        # Prepare the data for the KNN model
        X = np.array([[user_id, project_id] for other_user_id, project_id, _ in df_similar_projects])
        y = np.array([rating for _, _, rating in df_similar_projects])

        # Encode labels if necessary
        le = LabelEncoder()
        y = le.fit_transform(y)

        # Train the KNN model
        knn_model = KNeighborsClassifier(n_neighbors=k)
        knn_model.fit(X, y)

        # Predict the probability of liking each project
        projects_probabilities = knn_model.predict_proba(X)[:, 1]

        # Create a list of tuples (project, probability) and sort it in descending order of probability
        recommended_projects = list(zip(all_projects, projects_probabilities))
        recommended_projects.sort(key=lambda x: x[1], reverse=True)

        # Close the database connection
        conn.close()

        return recommended_projects[:10]  # Return top 10 recommended projects

    else:
        conn.close()
        return []

# API endpoint to get project ratings for a user
@app.route('/api/v1/model-ratings', methods=['GET'])
def get_ratings():
    user_id = get_id_from_header_request(request)
    if not user_id:
        return jsonify({'error': 'Valid user_id is required'}), 400
    
    conn = connect_to_db()
    cursor = conn.cursor()
    
    query = "SELECT id, interest FROM users WHERE id = %s;"
    cursor.execute(query, (user_id,))
    user_interest = cursor.fetchone()
    if user_interest:
        user_interest_name = user_interest[1]

        query = "SELECT id, interest FROM users WHERE interest = %s AND id != %s;"
        cursor.execute(query, (user_interest_name, user_id))
        users_with_same_interest = cursor.fetchall()

        query = "SELECT user_id, project_id, rating FROM ratings;"
        cursor.execute(query)
        ratings_data = cursor.fetchall()
        
        user_ratings = {}
        for rating in ratings_data:
            if rating[0] not in user_ratings:
                user_ratings[rating[0]] = {}
            user_ratings[rating[0]][rating[1]] = rating[2]

        recommended_projects = recommend_all_projects_prob(user_id, user_ratings)

        if not recommended_projects:
            # If no recommendations are made due to insufficient data, return a message and all projects in random order
            query = "SELECT id, name, description, macro_setor, micro_setor, image_link FROM projects;"
            cursor.execute(query)
            projects_data = cursor.fetchall()
            shuffle(projects_data)
            conn.close()
            return jsonify({'message': 'Insufficient data for recommendation. Returning all projects in random order.', 'projects': projects_data})

        conn.close()
        return jsonify(recommended_projects)

    else:
        conn.close()
        return jsonify({'error': 'User not found'}), 404

if __name__ == '__main__':
    app.run(debug=True)

