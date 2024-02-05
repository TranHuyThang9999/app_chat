import React, { useEffect, useState } from 'react'
import Home from './Home'
import { Button } from 'antd'
import Profile from './Profile'


export default function Routers() {

  const [showProfile, setShowProfile] = useState(false);
  let content = null;

  const handleLoginSuccess = () => {
    setShowProfile(true);
  };


  useEffect(() => {
    const user_name = localStorage.getItem('user_name');
    if (user_name) {
      setShowProfile(true);
    }
  }, []);

  const backToHome = () => {
    setShowProfile(false);  
  };
  
  if (showProfile){
    content = <Profile backToHome={backToHome} />;
  }else{
    content = <Home onLoginSuccess={handleLoginSuccess}/>
  }
  return (
    <div>
      {content}
    </div>
  )
}
