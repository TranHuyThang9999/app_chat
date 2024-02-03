import React, { useEffect, useState } from 'react';

export default function UserContext() {
  const [userName, setUserName] = useState('');

  useEffect(() => {
    const storedUserName = localStorage.getItem('user_name');
    if (storedUserName) {
      setUserName(storedUserName);
   //   localStorage.setItem('user_name',storedUserName);
    }
  }, []);
  if(!userName){
    return  null;
  }
  return userName;
}