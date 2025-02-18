import { useState } from 'react';
import { Form, Input, Button, Card, message } from 'antd';
import { UserOutlined, MailOutlined } from '@ant-design/icons';
import { Link, Navigate } from 'react-router-dom';
import axios from 'axios';

interface RegisterProps {
  isLoggedIn: boolean;
}

interface RegisterForm {
  username: string;
  email: string;
}

const Register: React.FC<RegisterProps> = ({ isLoggedIn }) => {
  const [loading, setLoading] = useState(false);

  if (isLoggedIn) {
    return <Navigate to="/users" replace />;
  }

  const onFinish = async (values: RegisterForm) => {
    setLoading(true);
    try {
      await axios.post('http://localhost:8000/api/v1/users', values);
      message.success('注册成功！请使用新账号登录。');
      window.location.href = '/login';
    } catch (error) {
      message.error('注册失败，请稍后重试！');
      console.error('Register error:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ 
      display: 'flex', 
      justifyContent: 'center', 
      alignItems: 'center', 
      height: '100vh',
      background: '#f0f2f5'
    }}>
      <Card title="用户注册" style={{ width: 400 }}>
        <Form
          name="register"
          onFinish={onFinish}
          autoComplete="off"
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: '请输入用户名！' }]}
          >
            <Input
              prefix={<UserOutlined />}
              placeholder="用户名"
            />
          </Form.Item>

          <Form.Item
            name="email"
            rules={[
              { required: true, message: '请输入邮箱！' },
              { type: 'email', message: '请输入有效的邮箱地址！' }
            ]}
          >
            <Input
              prefix={<MailOutlined />}
              placeholder="邮箱"
            />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading} block>
              注册
            </Button>
          </Form.Item>

          <div style={{ textAlign: 'center' }}>
            已有账号？<Link to="/login">立即登录</Link>
          </div>
        </Form>
      </Card>
    </div>
  );
};

export default Register;