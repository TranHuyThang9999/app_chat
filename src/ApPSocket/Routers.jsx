import React, { useEffect, useState } from 'react'
import Home from './Home'
import { Button } from 'antd'
import Profile from './Profile'


export default function Routers() {

  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [showProfile, setShowProfile] = useState(false);
  
  const handleLoginSuccess = () => {
    setIsLoggedIn(true);
    setShowProfile(true);
  };
  const handleLogout = () => {
    localStorage.removeItem('user_name');
    setIsLoggedIn(false);
    setShowProfile(false);
  };

  useEffect(() => {
    const user_name = localStorage.getItem('user_name');
    if (user_name) {
      setIsLoggedIn(true);
      setShowProfile(true);
    }
  }, []);

  return (
    <div>
      {isLoggedIn && (
        <Button onClick={handleLogout} style={{ marginBottom: '16px' }}>
          Logout
        </Button>
      )}

      {showProfile ? <Profile /> : <Home onLoginSuccess={handleLoginSuccess} />}
    </div>
  )
}
