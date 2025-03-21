export class UserService {
  #key = "@user_profile";

  async auth(code,provider) {
    try {
      const request = await fetch(`/api/v1/login/${code}?provider=${provider}`);
      const user = await request.json();

      localStorage.setItem(this.#key, JSON.stringify(user));

      return true;
    } catch (err) {
      console.error(err);
    }
  }

  getProfile() {
    const user = localStorage.getItem(this.#key);
    if (!user) {
      window.location.replace("/login");
    }

    return JSON.parse(user);
  }
}
