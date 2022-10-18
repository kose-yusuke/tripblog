import React from 'react';
import {BrowserRouter,Route,Routes} from 'react-router-dom'
import HomePage from './pages/Home'
import LoginPage from './pages/Login'
import UserInfo from './components/UserInfo'
import TodoDashboard from './pages/Dashboard'


const App = () => {
  
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage/>} />
        <Route path="/home" element={<HomePage/>} />
        <Route path="/user" element={<UserInfo/>} />
        <Route path="/todo" element={<TodoDashboard />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;