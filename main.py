# In[2]:


import requests
import json


# In[9]:


ADDRESS = "http://localhost:8080/data"

# Making a get request
response = requests.get(ADDRESS)

# Writing data in a json file
with open('data.json', 'w') as f:
    f.write(json.dumps(response.json(), indent=4))


# In[170]:


# Opening JSON file
f = open('data.json')


# In[171]:


# returns JSON object as 
# a dictionary
data = json.load(f)


# In[173]:


# python obj to json obj
routes = json.dumps(data['routes'])
searches = json.dumps(data['searches'])


# In[ ]:


import pandas as pd


# In[174]:


# store routes data into a csv file
df = pd.read_json(routes)
df.to_csv('data_routes.csv')


# In[1]:


# store searches data into a csv file
df = pd.read_json(searches)
df.to_csv('data_searches.csv')


# In[178]:


df_routes = pd.read_csv('data_routes.csv')



# In[181]:


# extract x, y from start and stop columns
import ast
df_routes['Xstart'] = df_routes.apply(lambda x: ast.literal_eval(x['start'])['x'], axis=1)


# In[182]:


df_routes['Ystart'] = df_routes.apply(lambda x: ast.literal_eval(x['start'])['y'], axis=1)


# In[183]:


df_routes['Xstop'] = df_routes.apply(lambda x: ast.literal_eval(x['stop'])['x'], axis=1)


# In[184]:


df_routes['Ystop'] = df_routes.apply(lambda x: ast.literal_eval(x['stop'])['y'], axis=1)


# In[186]:


# droping columns that aren't necessary
df_routes.drop(['Unnamed: 0', 'start', 'stop'], axis=1, inplace=True)


# In[187]:


df_searches = pd.read_csv('data_searches.csv')



# In[189]:


df_searches.drop(['Unnamed: 0'], axis=1, inplace=True)


# In[190]:


# finding closest station to what ppl searched
from shapely.geometry import Point, MultiPoint
from shapely.ops import nearest_points


# In[191]:


points = []

for i, row in df_routes.iterrows():
    point1 = Point(row['Xstart'], row['Ystart'])
    point2 = Point(row['Xstop'], row['Ystop'])
    points.append(point1)
    points.append(point2)


# In[192]:


destination = MultiPoint(points)


# In[193]:


df_searches['src'] = df_searches.apply(lambda x: nearest_points(Point(x['x1'],x['y1']), destination)[1], axis=1)
df_searches['dest'] = df_searches.apply(lambda x: nearest_points(Point(x['x2'],x['y2']), destination)[1], axis=1)



# In[196]:


df_searches.dest.value_counts()


# In[197]:


df_searches.src.value_counts()


# In[198]:


#creating paths
df_searches['path'] = df_searches['src'].map(str) + '-' + df_searches['dest'].map(str)



# In[200]:


# finding searched pathes that has benefit to add
print(df_searches.path.value_counts())
