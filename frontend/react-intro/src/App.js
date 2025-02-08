import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Table from './components/Table';
import './App.css';

function App() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true); 
  const [hasError, setHasError] = useState(false); 
  const [errorMessage, setErrorMessage] = useState(''); 

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:8086/api/containers');
        setData(response.data); 
        setHasError(false); 
      } catch (error) {
        console.error('get data error:', error);
        setHasError(true);
        setErrorMessage(error.message); 
      } finally {
        setLoading(false); 
      }
    };

    fetchData();
  }, []);

  return (
    <div className="App">
      <h1>Containers list</h1>

      {loading && <div>Loading...</div>}

      {hasError && (
        <div style={{ color: 'red' }}>
          Error happen: {errorMessage}
        </div>
      )}

      {!loading && !hasError && <Table data={data} />}
    </div>
  );
}

export default App;