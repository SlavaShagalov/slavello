import { createContext, PropsWithChildren, useContext, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import User from "../models/User";

import { AppDispatch, RootState } from "../services/state/store";
import { getByTokenAsync } from "../services/state/user/userSlice";

const AuthContext = createContext<User | null>(null);

type AuthProviderProps = PropsWithChildren & {
  isSignedIn?: boolean;
};

export default function AuthProvider({ children }: AuthProviderProps) {
  // const [isSignedIn, setIsSignedIn] = useState(false);

  const user = useSelector((state: RootState) => state.user.user);
  const dispatch = useDispatch<AppDispatch>();

  useEffect(() => {
    dispatch(getByTokenAsync());
  }, [dispatch]);

  return <AuthContext.Provider value={user}>{children}</AuthContext.Provider>;
}

export const useAuth = () => {
  const context = useContext(AuthContext);

  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }

  return context;
};
