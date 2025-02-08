import React from 'react';
import { Table } from 'antd';
import 'antd/dist/reset.css'; 

const TableComponent = ({ data }) => {
  const columns = [
    {
      title: 'Address',
      dataIndex: 'addr',
      key: 'addr',
    },
    {
      title: 'Last ping time',
      dataIndex: 'last_ping_time',
      key: 'last_ping_time',
      render: (text) => text || '-', 
    },
    {
      title: 'Last alive time',
      dataIndex: 'last_alive_time',
      key: 'last_alive_time',
      render: (text) => text || '-',
    },
  ];

  return (
    <Table
      columns={columns}
      dataSource={data}
      rowKey="addr" 
      pagination={false} 
    />
  );
};

export default TableComponent;