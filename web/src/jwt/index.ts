export function parse(token: string) {
  try {
    return JSON.parse(atob(token.split(".")[1]).toString());
  } catch (e) {
    return null;
  }
}
