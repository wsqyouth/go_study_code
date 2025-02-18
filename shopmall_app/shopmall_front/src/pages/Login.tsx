import { useState } from 'react';
import { Form, Input, Button, Card, message } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { Link, Navigate } from 'react-router-dom';
import axios from 'axios';

interface LoginProps {
  onLoginSuccess: (user: { username: string }) => void;
  isLoggedIn: boolean;
}

interface LoginForm {
  username: string;
  email: string;
}

const Login: React.FC<LoginProps> = ({ onLoginSuccess, isLoggedIn }) => {
  const [loading, setLoading] = useState(false);

  if (isLoggedIn) {
    return <Navigate to="/users" replace />;
  }

  const onFinish = async (values: LoginForm) => {
    setLoading(true);
    try {
      const response = await axios.post('http://localhost:8000/api/v1/login', values);
      const user = response.data;
      message.success(`欢迎回来，${user.username}！`);
      onLoginSuccess({ username: user.username });
    } catch (error) {
      message.error('登录失败，请稍后重试！');
      console.error('Login error:', error);
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
      <Card title="用户登录" style={{ width: 400 }}>
        <Form
          name="login"
          initialValues={{ remember: true }}
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
            rules={[{ required: true, message: '请输入邮箱！' }]}
          >
            <Input
              prefix={<LockOutlined />}
              placeholder="邮箱"
            />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading} block>
              登录
            </Button>
          </Form.Item>

          <div style={{ textAlign: 'center' }}>
            还没有账号？<Link to="/register">立即注册</Link>
          </div>
        </Form>
      </Card>
    </div>
  );
};

export default Login;