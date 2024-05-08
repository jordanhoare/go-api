// src/frontend/src/index.tsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './pages/app/App';
import NotFoundPage from './pages/NotFound';
import DataPage from './pages/hello/Hello';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/hello" element={<DataPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>,
);
