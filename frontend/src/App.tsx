import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import SignUpPage from "./pages/SignUpPage";
import SignInPage from "./pages/SignInPage";
import WorkspacesPage from "./pages/WorkspacesPage/WorkspacesPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signin" element={<SignInPage />} />
        <Route path="/signup" element={<SignUpPage />} />
        <Route path="/workspaces" element={<WorkspacesPage />} />
        <Route path="/" element={<WorkspacesPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
