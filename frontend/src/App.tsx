import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import SignUpPage from "./pages/SignUpPage";
import SignInPage from "./pages/SignInPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signin" element={<SignInPage />} />
        <Route path="/signup" element={<SignUpPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
