// POST /api/waitlist: a Coder waiting-list sign-up from the form on commonprefix.com/coder.
// Adds a row to the "Coder waiting list" database in Common Prefix's Notion (Name, Email;
// Date is set by Notion; Status starts at "Follow up").
// Needs NOTION_TOKEN in the Vercel project's environment variables: the secret of a Notion
// internal integration that has been given access to that database.

const DATABASE_ID = "422bf672de82816a9728e7d616c0223a";
const EMAIL = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

module.exports = async function handler(req, res) {
  if (req.method !== "POST") {
    res.setHeader("Allow", "POST");
    return res.status(405).json({ ok: false, error: "method" });
  }
  let body = req.body;
  if (typeof body === "string") {
    try {
      body = JSON.parse(body);
    } catch {
      body = {};
    }
  }
  body = body || {};
  // A field people never see: bots fill it in, people do not.
  if (body.company) return res.status(200).json({ ok: true });

  const name = String(body.name || "").trim().slice(0, 120);
  const email = String(body.email || "").trim().slice(0, 200);
  if (!name || !EMAIL.test(email)) return res.status(400).json({ ok: false, error: "invalid" });

  const token = process.env.NOTION_TOKEN;
  if (!token) return res.status(503).json({ ok: false, error: "not-configured" });

  const r = await fetch("https://api.notion.com/v1/pages", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Notion-Version": "2022-06-28",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      parent: { database_id: DATABASE_ID },
      properties: {
        Name: { title: [{ text: { content: name } }] },
        Email: { email },
        Status: { status: { name: "Follow up" } },
      },
    }),
  });
  if (!r.ok) {
    console.error("Notion", r.status, await r.text());
    return res.status(502).json({ ok: false, error: "store" });
  }
  return res.status(200).json({ ok: true });
};
