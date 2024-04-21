import React from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

import AuthProvider from "./components/AuthProvider";
import ProtectedRoute from "./components/ProtectedRoute";

import {
  BOARD_PAGE_URL,
  SIGNIN_PAGE_URL,
  SIGNUP_PAGE_URL,
  WORKSPACES_PAGE_URL,
} from "./constants";

import SignUpPage from "./pages/SignUpPage";
import SignInPage from "./pages/SignInPage";
import WorkspacesPage from "./pages/WorkspacesPage/WorkspacesPage";
import BoardPage from "./pages/BoardPage/BoardPage";

const router = createBrowserRouter([
  {
    path: SIGNUP_PAGE_URL,
    element: <SignUpPage />,
  },
  {
    path: SIGNIN_PAGE_URL,
    element: <SignInPage />,
  },
  {
    path: WORKSPACES_PAGE_URL,
    element: (
      <ProtectedRoute>
        <WorkspacesPage />
      </ProtectedRoute>
    ),
  },
  {
    path: BOARD_PAGE_URL,
    element: (
      <ProtectedRoute>
        <BoardPage />
      </ProtectedRoute>
    ),
  },
  {
    path: "/",
    element: (
      <ProtectedRoute>
        <WorkspacesPage />
      </ProtectedRoute>
    ),
  },
]);

function App() {
  return (
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  );
}

export default App;
