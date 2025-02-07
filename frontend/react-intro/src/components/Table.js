import React from 'react';

const Table = ({ data }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Address</th>
          <th>Last ping time</th>
          <th>Last alive time</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item, index) => (
          <tr key={index}>
            <td>{item.addr}</td>
            <td>{item.last_ping_time || ''}</td>
            <td>{item.last_alive_time || ''}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default Table;