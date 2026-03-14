const nodeData = {
  overview: {
    title: "System overview",
    summary:
      "The main request flow is: user enters the domain, Fightdb.org resolves through Cloudflare, Cloudflare forwards to Caddy on the Hetzner VPS, and Caddy proxies into the Go back-end. Inside the app stack, the Go service serves the front-end and reads from PostgreSQL, while Python retrieval jobs feed fresh data into the database.",
    details: [
      "Request flow: User -> Fightdb.org -> Cloudflare -> Caddy -> Go back-end.",
      "Presentation layer: Front-end uses JavaScript, HTML, and CSS.",
      "Application layer: Go handles routing and data access.",
      "Persistence layer: PostgreSQL stores the site data.",
      "Ingestion flow: Python retrieval jobs scrape and write data into PostgreSQL."
    ],
    tags: ["Portfolio project", "Full stack", "Self-hosted"],
    status: "Overview mode"
  },
  user: {
    title: "User",
    summary:
      "The end user enters the site URL into a browser. That browser request kicks off the rest of the path through DNS, the edge, the server, and the application.",
    details: [
      "The browser is the first client in the chain.",
    ],
    tags: ["Browser", "Client", "Request origin"],
    status: "Focused on user entry"
  },
  domain: {
    title: "FightDB.org",
    summary:
      "The domain is the stable public address for the project. It gives the site a clear identity and is the piece users actually remember and type.",
    details: [
      "The domain points traffic toward Cloudflare first.",
      "It separates the public-facing URL from the underlying server IP.",
    ],
    tags: ["Domain name", "Public endpoint", "Branding"],
    status: "Focused on the public entry point"
  },
  cloudflare: {
    title: "Cloudflare",
    summary:
      "Cloudflare resolves DNS and forwards requests to the origin server. It sits between the public internet and the VPS so the site has a cleaner, more resilient edge layer.",
    details: [
      "Acts as the DNS layer for FightDB.org.",
      "Forwards requests to the origin instead of exposing the app directly.",
      "Adds a cleaner separation between public traffic and the application host."
    ],
    tags: ["DNS", "Proxy", "Edge layer"],
    status: "Focused on DNS and the edge layer"
  },
  caddy: {
    title: "Caddy Server on the Hetzner VPS",
    summary:
      "Caddy runs on the Ubuntu VPS, terminates HTTPS, and reverse proxies requests into the Go app. That keeps TLS and web-server concerns separate from the application code.",
    details: [
      "Handles HTTPS termination before traffic reaches the app.",
      "Reverse proxies inbound requests to the Go process running internally.",
      "Sits on the Hetzner Ubuntu VPS as the public-facing web server."
    ],
    tags: ["Caddy", "HTTPS", "Reverse proxy", "Hetzner VPS", "Ubuntu"],
    status: "Focused on the VPS entry layer"
  },
  vps: {
    title: "Hetzner VPS (Ubuntu)",
    summary:
      "This is the deployed host for the application stack. The server boundary groups the Caddy layer, the Go app, the front-end assets, the Postgres database, and the supporting retrieval jobs.",
    details: [
      "The VPS is the origin Cloudflare forwards requests to.",
      "Ubuntu provides the runtime environment for the deployed services.",
      "Showing the VPS boundary makes it clear which parts live on the server and which parts are outside it."
    ],
    tags: ["Hetzner", "Ubuntu", "Origin server", "Hosting boundary"],
    status: "Focused on the hosting boundary"
  },
  frontend: {
    title: "Front-end",
    summary:
      "The front-end is the browser-facing presentation layer built with standard web technologies. It is the part the user ultimately sees after the request moves through the server-side stack.",
    details: [
      "Uses JavaScript for client-side behavior.",
      "Uses HTML for structure.",
      "Uses CSS for presentation and layout.",
      "In this flow, the Go app is responsible for serving the front-end to the browser."
    ],
    tags: ["JavaScript", "HTML", "CSS", "Presentation layer"],
    status: "Focused on the front-end"
  },
  backend: {
    title: "Go back-end",
    summary:
      "The Go service handles routing, page rendering, and database access. The implementation leans on stock packages where possible to keep the stack understandable and lightweight.",
    details: [
      "Uses net/http for routing and request handling.",
      "Uses database/sql with pgx for efficient PostgreSQL access.",
      "Choosing standard packages reduces dependency weight and makes the code path easier to reason about in a smaller personal project.",
      "A thinner dependency surface also makes deployment and long-term maintenance simpler."
    ],
    tags: ["Go", "net/http", "database/sql", "pgx", "Minimal dependencies"],
    status: "Focused on application logic"
  },
  database: {
    title: "PostgreSQL database",
    summary:
      "Postgres stores the structured data the site serves: fighters, matchups, events, and related records. It is the source of data for the user-facing application.",
    details: [
      "Back-end requests read from Postgres to build site pages.",
      "Python retrieval jobs write normalized data into the same database.",
      "A relational database fits this project well because fights, events, and fighters all have clear relationships."
    ],
    tags: ["PostgreSQL", "Source of truth", "Relational data"],
    status: "Focused on persistence"
  },
  scrapers: {
    title: "Python data retrieval",
    summary:
      "Python handles the data ingestion side of the system. Separate scripts scrape MMA sources, transform the data into the shape the app expects, and insert it into Postgres.",
    details: [
      "BeautifulSoup 4 is used for HTML scraping and extraction.",
      "psycopg2 is used as the PostgreSQL adapter for writes and updates.",
      "Keeping scraping separate from the Go app prevents the request-serving path from being mixed with ingestion jobs."
    ],
    tags: ["Python", "BeautifulSoup 4", "psycopg2", "ETL"],
    status: "Focused on data ingestion"
  }
};

const scaleFocused = 1.1;
const stageWidth = 1280;
const stageHeight = 920;

const panel = document.getElementById("detailsPanel");
const stage = document.getElementById("diagramStage");
const viewport = document.getElementById("diagramViewport");
const statusLabel = document.getElementById("diagramStatus");
const resetButton = document.getElementById("resetView");
const nodes = Array.from(document.querySelectorAll("button.node[data-node]"));
const viewState = {
  scale: 1,
  x: 0,
  y: 0,
  activeKey: "overview"
};
const dragState = {
  active: false,
  moved: false,
  pointerId: null,
  startX: 0,
  startY: 0,
  originX: 0,
  originY: 0
};

function clampPosition(x, y, scale) {
  const viewportWidth = viewport.clientWidth;
  const viewportHeight = viewport.clientHeight;
  const scaledWidth = stageWidth * scale;
  const scaledHeight = stageHeight * scale;

  let minX = viewportWidth - scaledWidth;
  let maxX = 0;
  let minY = viewportHeight - scaledHeight;
  let maxY = 0;

  if (scaledWidth <= viewportWidth) {
    minX = maxX = (viewportWidth - scaledWidth) / 2;
  }

  if (scaledHeight <= viewportHeight) {
    minY = maxY = (viewportHeight - scaledHeight) / 2;
  }

  return {
    x: Math.min(maxX, Math.max(minX, x)),
    y: Math.min(maxY, Math.max(minY, y))
  };
}

function updateStageTransform() {
  stage.style.transform = `translate(${viewState.x}px, ${viewState.y}px) scale(${viewState.scale})`;
}

function getOverviewScale() {
  const viewportWidth = viewport.clientWidth;
  const viewportHeight = viewport.clientHeight;
  const widthScale = viewportWidth / stageWidth;
  const heightScale = viewportHeight / stageHeight;

  return Math.min(widthScale, heightScale) * 0.94;
}

function setView(x, y, scale) {
  const next = clampPosition(x, y, scale);
  viewState.x = next.x;
  viewState.y = next.y;
  viewState.scale = scale;
  updateStageTransform();
}

function renderPanel(key) {
  const info = nodeData[key];
  const detailsMarkup = info.details
    .map((detail) => `<li>${detail}</li>`)
    .join("");
  const tagsMarkup = info.tags
    .map((tag) => `<span class="detail-tag">${tag}</span>`)
    .join("");

  panel.innerHTML = `
    <h2>${info.title}</h2>
    <p>${info.summary}</p>
    <div class="detail-tag-row">${tagsMarkup}</div>
    <h3>What this part is doing</h3>
    <ul class="detail-list">${detailsMarkup}</ul>
  `;

  statusLabel.textContent = info.status;
}

function setActiveNode(key) {
  nodes.forEach((node) => {
    node.classList.toggle("active", node.dataset.node === key);
  });
}

function focusNode(key) {
  const node = nodes.find((item) => item.dataset.node === key);
  if (!node) {
    resetView();
    return;
  }

  const x = Number(node.dataset.x);
  const y = Number(node.dataset.y);
  const viewportWidth = viewport.clientWidth;
  const viewportHeight = viewport.clientHeight;
  const translateX = viewportWidth / 2 - x * scaleFocused;
  const translateY = viewportHeight / 2 - y * scaleFocused;

  setView(translateX, translateY, scaleFocused);
  renderPanel(key);
  setActiveNode(key);
  viewState.activeKey = key;
}

function resetView() {
  const overviewScale = getOverviewScale();
  const viewportWidth = viewport.clientWidth;
  const viewportHeight = viewport.clientHeight;
  const translateX = (viewportWidth - stageWidth * overviewScale) / 2;
  const translateY = (viewportHeight - stageHeight * overviewScale) / 2;

  setView(translateX, translateY, overviewScale);
  renderPanel("overview");
  setActiveNode("");
  viewState.activeKey = "overview";
}

function handlePointerDown(event) {
  if (event.button !== 0) {
    return;
  }

  if (event.target.closest("button.node")) {
    return;
  }

  dragState.active = true;
  dragState.moved = false;
  dragState.pointerId = event.pointerId;
  dragState.startX = event.clientX;
  dragState.startY = event.clientY;
  dragState.originX = viewState.x;
  dragState.originY = viewState.y;
  viewport.classList.add("dragging");
  viewport.setPointerCapture(event.pointerId);
}

function handlePointerMove(event) {
  if (!dragState.active || event.pointerId !== dragState.pointerId) {
    return;
  }

  const deltaX = event.clientX - dragState.startX;
  const deltaY = event.clientY - dragState.startY;

  if (Math.abs(deltaX) > 4 || Math.abs(deltaY) > 4) {
    dragState.moved = true;
  }

  setView(
    dragState.originX + deltaX,
    dragState.originY + deltaY,
    viewState.scale
  );
}

function endDrag(event) {
  if (!dragState.active || event.pointerId !== dragState.pointerId) {
    return;
  }

  dragState.active = false;
  viewport.classList.remove("dragging");

  if (viewport.hasPointerCapture(event.pointerId)) {
    viewport.releasePointerCapture(event.pointerId);
  }
}

function suppressClickAfterDrag(event) {
  if (!dragState.moved) {
    return;
  }

  event.preventDefault();
  event.stopPropagation();
  dragState.moved = false;
}

nodes.forEach((node) => {
  node.addEventListener("click", () => focusNode(node.dataset.node));
});

viewport.addEventListener("pointerdown", handlePointerDown);
viewport.addEventListener("pointermove", handlePointerMove);
viewport.addEventListener("pointerup", endDrag);
viewport.addEventListener("pointercancel", endDrag);
viewport.addEventListener("click", suppressClickAfterDrag, true);

resetButton.addEventListener("click", resetView);
window.addEventListener("resize", () => {
  if (viewState.activeKey === "overview") {
    resetView();
    return;
  }

  const current = nodes.find((node) => node.dataset.node === viewState.activeKey);
  if (current) {
    focusNode(current.dataset.node);
  }
});

resetView();
