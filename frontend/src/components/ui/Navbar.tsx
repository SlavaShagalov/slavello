import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { API_HOST } from '../../constants';

const Navbar: React.FC = () => {
  const defaultData = {
    "id": 0,
    "username": "",
    "email": "",
    "name": "",
    "avatar": "",
    "created_at": "",
    "updated_at": ""
  };
  const [data, setData] = React.useState(defaultData);

  useEffect(() => {
    fetch(`${API_HOST}/api/v1/auth/me`, { credentials: 'include' })
      .then(response => {
        console.log("Status:", response.status);
        if (response.status === 200) {
          console.log('ws success');
        } else {
          console.log('ws failed');
        }
        return response.json();
      })
      .then(resultJson => {
        console.log(resultJson);
        console.log('parse success');

        setData(resultJson);
      })
      .catch(error => {
        console.log('error', error);
      });
  }, []);

  return (
    <div className="bg-green-400 h-12 p-2 flex items-center justify-between">
      <Link to={"/workspaces"}>
        <img src="/assets/Logo.svg" alt="Logo" className="rounded-lg"/>
      </Link>
      <Link to={"/settings"}>
        {/* <img className='w-8 h-8' src={data.avatar} alt="Settings" /> */}
        <img className='w-8 h-8' src="/assets/Avatar.png" alt="Settings" />
      </Link>
    </div >
  );
}

export default Navbar;
