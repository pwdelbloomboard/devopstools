# Import dependencies
import pandas as pd
import numpy as np
import sklearn
import joblib

def load_datacamp_trainingcsv(url):
    # Load the dataset in a dataframe object and include only four features as mentioned
    df = pd.read_csv(url)
    include = ['Age', 'Sex', 'Embarked', 'Survived'] # Only four features
    df_ = df[include]
    print(df_.head())
    return df_

def replace_nans_with_zeros(input_csv_df):

    categoricals = []
    for col, col_type in input_csv_df.dtypes.items():
        if col_type == 'O':
            categoricals.append(col)
        else:
            input_csv_df[col].fillna(0, inplace=True)

    print(categoricals)

def encoding_and_logistic_regression(categoricals,df_):
    # Encode categorical features as a one-hot numeric array.
    df_ohe = pd.get_dummies(df_, columns=categoricals, dummy_na=True)
    
    from sklearn.linear_model import LogisticRegression
    dependent_variable = 'Survived'
    x = df_ohe[df_ohe.columns.difference([dependent_variable])]
    y = df_ohe[dependent_variable]
    lr = LogisticRegression()
    lr.fit(x, y)
    
    LogisticRegression(C=1.0, class_weight=None, dual=False, fit_intercept=True,
    intercept_scaling=1, max_iter=100, multi_class='ovr', n_jobs=1,
    penalty='l2', random_state=None, solver='liblinear', tol=0.0001,
    verbose=0, warm_start=False)
    
    print(lr)
    
    return(lr)

def save_model_file_to_tmp(lr):
    FILENAME = '/tmp/model.joblib'
    joblib.dump(lr, FILENAME)
    print('File should be available at /tmp/model.joblib')

def execute_create_and_save_joblib():
    INPUT_URL = "http://s3.amazonaws.com/assets.datacamp.com/course/Kaggle/train.csv"
    kaggle_trainingcsv_df = load_datacamp_trainingcsv(INPUT_URL)
    kaggle_categories = replace_nans_with_zeros(kaggle_trainingcsv_df)
    kaggle_lr = encoding_and_logistic_regression(kaggle_categories,kaggle_trainingcsv_df)
    save_model_file_to_tmp(kaggle_lr)
