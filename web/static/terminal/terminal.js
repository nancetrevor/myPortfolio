const go = new Go();
let wasmReady = false;

async function initWasm() {
  try {
    const result = await WebAssembly.instantiateStreaming(
      fetch("/static/terminal/main.wasm"),
      go.importObject
    );
    go.run(result.instance);
    wasmReady = true;
  } catch (err) {
    console.error("Failed to load Wasm:", err);
  }
}

initWasm();

const terminal = document.getElementById("terminal");

let history = ['Type "help" and press Enter.'];
let currentInput = "";

function renderTerminal() {
  const lines = [...history, `> ${currentInput}<span class="cursor">█</span>`];
  terminal.innerHTML = lines.join("\n");
  terminal.scrollTop = terminal.scrollHeight;
}

terminal.addEventListener("click", () => {
  terminal.focus();
});

document.addEventListener("keydown", (e) => {
  if (document.activeElement !== terminal) return;

  if (e.key === "Backspace") {
    e.preventDefault();
    currentInput = currentInput.slice(0, -1);
    renderTerminal();
    return;
  }

  if (e.key === "Enter") {
    e.preventDefault();

    const command = currentInput;
    history.push(`> ${command}`);

    if (!wasmReady || typeof window.runCommand !== "function") {
      history.push("Wasm is not loaded yet.");
    } else {
      const output = window.runCommand(command);

      if (output === "__CLEAR__") {
        history = [];
      } else if (output.startsWith("__NAVIGATE__:")) {
        const url = output.replace("__NAVIGATE__:", "");
        history.push("Opening fightdb.org...")
        setTimeout(() => window.open(url, "_blank"), 600);
      }
      else if (output) {
        history.push(output);
      }
    }

    currentInput = "";
    renderTerminal();
    return;
  }

  if (e.key.length === 1 && !e.ctrlKey && !e.metaKey && !e.altKey) {
    e.preventDefault();
    currentInput += e.key;
    renderTerminal();
  }
});

renderTerminal();
terminal.focus();
