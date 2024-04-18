import { API_HOST } from "../constants";

const API_WORKSPACES_URL = API_HOST + "/api/v1/workspaces";

class WorkspaceService {
  static async create() {
    const requestOptions: RequestInit = {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: "New workspace",
      }),
    };

    const response = await fetch(
      API_WORKSPACES_URL,
      requestOptions
    );
    if (response.status === 200) {
      return await response.json();
    } else {
      throw new Error("Failed to create workspace");
    }
  }

  static async list() {
    const response = await fetch(API_WORKSPACES_URL, {
      credentials: "include",
    });
    if (response.status === 200) {
      return await response.json();
    } else {
      throw new Error("Failed to create board");
    }
  }

  static async delete(id: number) {
    const requestOptions: RequestInit = {
      method: "DELETE",
      credentials: "include",
    };

    const response = await fetch(
      `${API_WORKSPACES_URL}/${id}`,
      requestOptions
    );
    if (response.status === 204) {
      return true;
    } else {
      throw new Error("Failed to delete list");
    }
  }

  static async updateName(id: number, newName: string) {
    const requestOptions: RequestInit = {
      method: "PATCH",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: newName,
      }),
    };

    const response = await fetch(
      `${API_WORKSPACES_URL}/${id}`,
      requestOptions
    );
    if (response.status === 200) {
      return response.json();
    } else {
      throw new Error("Failed to update list name");
    }
  }
}

export default WorkspaceService;
