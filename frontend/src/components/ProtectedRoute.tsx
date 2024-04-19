import { useNavigate } from "react-router-dom";
import { useSelector } from "react-redux";

import { RootState } from "../services/state/store";

import Spinner from "./ui/Spinner";

const ProtectedRoute: React.FC<{
  children: any;
}> = ({ children }) => {
  const navigate = useNavigate();

  const user = useSelector((state: RootState) => state.user.user);
  const userStatus = useSelector((state: RootState) => state.user.status);

  let content;
  console.log("ProtectedRoute:", "status=", userStatus, "user=", user);
  if (userStatus === "loading") {
    content = (
      <div className="flex items-center justify-center h-screen">
        <Spinner />
      </div>
    );
  } else if (userStatus === "succeeded" && user !== null) {
    content = children;
  } else if (userStatus === "failed" || user === null) {
    navigate("/signin");
    content = (
      <div className="flex items-center justify-center h-full">
        <Spinner />
      </div>
    );
  } else {
    content = (
      <div className="flex items-center justify-center h-full">
        <Spinner />
      </div>
    );
  }

  return content;
};

export default ProtectedRoute;
