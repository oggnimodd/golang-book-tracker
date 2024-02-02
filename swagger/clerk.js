const header = document.querySelector(".topbar-wrapper");
const loginButton = document.createElement("button");
const logoutButton = document.createElement("button");

loginButton.innerText = "Sign In";
logoutButton.innerText = "Sign Out";

const authButtonWrapper = document.createElement("div");
loginButton.onclick = () =>
  Clerk.openSignIn({
    redirectUrl: window.location.href,
  });
logoutButton.onclick = () => Clerk.signOut();

// Wrap buttons
authButtonWrapper.appendChild(loginButton);
authButtonWrapper.appendChild(logoutButton);

// Add it to the header
header.appendChild(authButtonWrapper);

const startClerk = async () => {
  const Clerk = window.Clerk;

  try {
    await Clerk.load();
  } catch (err) {
    console.error("Error starting Clerk: ", err);
  }
};

(async () => {
  // Getting the CLERK_PUBLISHABLE_KEY from server
  const { CLERK_PUBLISHABLE_KEY } = await fetch("/CLERK_PUBLISHABLE_KEY").then(
    (res) => res.json(),
  );

  const script = document.createElement("script");
  script.setAttribute("data-clerk-publishable-key", CLERK_PUBLISHABLE_KEY);
  script.async = true;
  script.src =
    "https://cdn.jsdelivr.net/npm/@clerk/clerk-js@latest/dist/clerk.browser.js";
  script.crossOrigin = "anonymous";
  script.addEventListener("load", startClerk);
  script.addEventListener("error", () => {
    document.getElementById("no-frontend-api-warning").hidden = false;
  });
  document.body.appendChild(script);
})();
