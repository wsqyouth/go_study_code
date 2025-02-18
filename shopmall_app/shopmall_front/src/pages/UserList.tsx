import { useState, useEffect } from 'react';
import { Table, Button, Space, Modal, Form, Input, message, Layout, Typography } from 'antd';
import { EditOutlined, DeleteOutlined, LogoutOutlined, UserOutlined } from '@ant-design/icons';
import axios from 'axios';

interface User {
  id: number;
  username: string;
  email: string;
}

interface UserListProps {
  currentUser: { username: string } | null;
  onLogout: () => void;
}

const { Header, Content } = Layout;
const { Title } = Typography;

const UserList: React.FC<UserListProps> = ({ currentUser, onLogout }) => {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [editingUser, setEditingUser] = useState<User | null>(null);
  const [form] = Form.useForm();

  const fetchUsers = async () => {
    setLoading(true);
    try {
      const response = await axios.get('http://localhost:8000/api/v1/users?page=1&page_size=10');
      setUsers(response.data.users || []);
    } catch (error) {
      message.error({
        content: '获取用户列表失败！',
        style: {
          marginTop: '20vh',
        },
      });
      console.error('Fetch users error:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleEdit = (user: User) => {
    if (currentUser?.username !== 'root') {
      message.error({
        content: '只有 root 用户才能编辑用户信息！',
        style: {
          marginTop: '20vh',
        },
      });
      return;
    }
    if (user.username === 'root') {
      message.error({
        content: '不能编辑 root 用户！',
        style: {
          marginTop: '20vh',
        },
      });
      return;
    }
    setEditingUser(user);
    form.setFieldsValue(user);
    setEditModalVisible(true);
  };

  const handleDelete = async (user: User) => {
    if (currentUser?.username !== 'root') {
      message.error({
        content: '只有 root 用户才能删除用户！',
        style: {
          marginTop: '20vh',
        },
      });
      return;
    }
    if (user.username === 'root') {
      message.error({
        content: '不能删除 root 用户！',
        style: {
          marginTop: '20vh',
        },
      });
      return;
    }

    Modal.confirm({
      title: '确认删除',
      content: `确定要删除用户 ${user.username} 吗？`,
      onOk: async () => {
        try {
          await axios.delete(`http://localhost:8000/api/v1/users/${user.id}`);
          message.success({
            content: '删除成功！',
            style: {
              marginTop: '20vh',
            },
          });
          fetchUsers();
        } catch (error) {
          message.error({
            content: '删除失败！',
            style: {
              marginTop: '20vh',
            },
          });
          console.error('Delete user error:', error);
        }
      },
    });
  };

  const handleEditSubmit = async () => {
    if (currentUser?.username !== 'root') {
      message.error({
        content: '只有 root 用户才能编辑用户信息！',
        style: {
          marginTop: '20vh',
        },
      });
      return;
    }

    try {
      const values = await form.validateFields();
      if (editingUser) {
        await axios.put(`http://localhost:8000/api/v1/users/${editingUser.id}`, values);
        message.success({
          content: '更新成功！',
          style: {
            marginTop: '20vh',
          },
        });
        setEditModalVisible(false);
        fetchUsers();
      }
    } catch (error) {
      message.error({
        content: '更新失败！',
        style: {
          marginTop: '20vh',
        },
      });
      console.error('Update user error:', error);
    }
  };

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: '10%',
      align: 'center' as const
    },
    {
      title: '用户名',
      dataIndex: 'username',
      key: 'username',
      width: '35%',
      align: 'center' as const
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      key: 'email',
      width: '35%',
      align: 'center' as const
    },
    {
      title: '操作',
      key: 'action',
      width: '20%',
      align: 'center' as const,
      render: (_: any, record: User) => (
        <Space size="middle" style={{ display: 'flex', justifyContent: 'center' }}>
          <Button
            type="primary"
            icon={<EditOutlined />}
            onClick={() => handleEdit(record)}
            disabled={currentUser?.username !== 'root'}
            size="small"
            style={{
              borderRadius: '4px',
              display: 'flex',
              alignItems: 'center',
              gap: '4px'
            }}
          >
            编辑
          </Button>
          <Button
            danger
            icon={<DeleteOutlined />}
            onClick={() => handleDelete(record)}
            disabled={currentUser?.username !== 'root'}
            size="small"
            style={{
              borderRadius: '4px',
              display: 'flex',
              alignItems: 'center',
              gap: '4px'
            }}
          >
            删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh', background: '#f5f7fa' }}>
      <Header style={{ 
        padding: '0 32px', 
        background: '#fff', 
        borderBottom: '1px solid #e8e8e8', 
        display: 'flex', 
        justifyContent: 'space-between', 
        alignItems: 'center',
        boxShadow: '0 2px 8px rgba(0,0,0,0.06)'
      }}>
        <Title level={4} style={{ margin: 0, color: '#1890ff', display: 'flex', alignItems: 'center', gap: '8px' }}>
          <UserOutlined style={{ fontSize: '24px' }} /> 用户管理系统
        </Title>
        <Space>
          <span style={{ color: '#666', fontSize: '14px' }}>欢迎回来，{currentUser?.username}！</span>
          <Button 
            type="primary" 
            icon={<LogoutOutlined />} 
            onClick={onLogout}
            style={{
              background: '#1890ff',
              borderColor: '#1890ff',
              boxShadow: '0 2px 0 rgba(24,144,255,0.1)',
              display: 'flex',
              alignItems: 'center',
              gap: '4px',
              borderRadius: '4px'
            }}
          >
            退出登录
          </Button>
        </Space>
      </Header>
      <Content style={{ padding: '32px', background: '#f5f7fa' }}>
        <div style={{ 
          background: '#fff', 
          padding: '24px', 
          borderRadius: '12px', 
          boxShadow: '0 1px 2px rgba(0, 0, 0, 0.03)',
          border: '1px solid #f0f0f0'
        }}>
          <Table
            columns={columns}
            dataSource={users}
            rowKey="id"
            loading={loading}
            style={{ 
              '.ant-table': { fontSize: '14px' },
              '.ant-table-thead > tr > th': {
                background: '#fafafa',
                fontWeight: 600,
                color: '#1f1f1f'
              },
              '.ant-table-tbody > tr:hover > td': {
                background: '#f5f7fa'
              }
            }}
            pagination={{
              showSizeChanger: true,
              showQuickJumper: true,
              showTotal: (total) => `共 ${total} 条记录`,
              pageSize: 10,
              style: {
                marginTop: '16px',
                textAlign: 'center',
                '.ant-pagination-item-active': {
                  borderColor: '#1890ff',
                  color: '#1890ff'
                }
              }
            }}
          />
        </div>

        <Modal
          title={<div style={{ color: '#1890ff', fontWeight: 600 }}>编辑用户</div>}
          open={editModalVisible}
          onOk={handleEditSubmit}
          onCancel={() => setEditModalVisible(false)}
          maskClosable={false}
          destroyOnClose
          style={{ '.ant-modal-content': { borderRadius: '12px' } }}
        >
          <Form form={form} layout="vertical" style={{ '.ant-form-item-label > label': { color: '#1f1f1f' } }}>
            <Form.Item
              name="username"
              label="用户名"
              rules={[{ required: true, message: '请输入用户名！' }]}
            >
              <Input style={{ borderRadius: '6px' }} />
            </Form.Item>
            <Form.Item
              name="email"
              label="邮箱"
              rules={[
                { required: true, message: '请输入邮箱！' },
                { type: 'email', message: '请输入有效的邮箱地址！' }
              ]}
            >
              <Input style={{ borderRadius: '6px' }} />
            </Form.Item>
          </Form>
        </Modal>
      </Content>
    </Layout>
  );
};

export default UserList;