type User = {
  id: bigint;
  username: string;
  email: string;
  name: string;
  avatar: string | null;
  created_at: Date;
  updated_at: Date;
};

export default User;
