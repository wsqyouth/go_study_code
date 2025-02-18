import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import { useState } from 'react'
import Register from './pages/Register'
import Login from './pages/Login'
import UserList from './pages/UserList'
import './App.css'

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentUser, setCurrentUser] = useState<{ username: string } | null>(null);

  const handleLoginSuccess = (user: { username: string }) => {
    setIsLoggedIn(true);
    setCurrentUser(user);
  };

  return (
    <Router>
      <Routes>
        <Route path="/register" element={<Register isLoggedIn={isLoggedIn} />} />
        <Route path="/login" element={<Login onLoginSuccess={handleLoginSuccess} isLoggedIn={isLoggedIn} />} />
        <Route path="/users" element={isLoggedIn ? <UserList currentUser={currentUser} /> : <Navigate to="/login" />} />
        <Route path="/" element={<Navigate to="/login" replace />} />
      </Routes>
    </Router>
  )
}

export default App
