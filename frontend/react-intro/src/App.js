import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Table from './components/Table';
import './App.css';

function App() {
  const [data, setData] = useState([]);

  useEffect(() => {
    // Функция для получения данных с API
    const fetchData = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:8086/api/containers');
        setData(response.data);
      } catch (error) {
        console.error('Ошибка при получении данных:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="App">
      <h1>Containers list</h1>
      <Table data={data} />
    </div>
  );
}

export default App;