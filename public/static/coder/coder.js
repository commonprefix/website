// Common Prefix Coder page: reveal-on-scroll and the progress rings.
;(function () {
	const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches
	const ease = (x) => (x < 0.5 ? 4 * x * x * x : 1 - Math.pow(-2 * x + 2, 3) / 2)
	const C = 2 * Math.PI * 80

	function setArc(el, from, len) {
		el.style.strokeDasharray = `${len * C} ${C}`
		el.style.strokeDashoffset = `${-from * C}`
	}

	function playRing(ring) {
		const target = +ring.dataset.ring / 100
		const fill = ring.querySelector('.cx-ring__fill')
		const gap = ring.querySelector('.cx-ring__gap')
		const count = ring.querySelector('.cx-ring__count')
		const dur = reduced ? 0 : 1800
		const t0 = performance.now()
		const step = (now) => {
			const x = dur ? Math.min(1, (now - t0) / dur) : 1
			const v = target * ease(x)
			setArc(fill, 0, v)
			count.textContent = Math.round(v * 100) + '%'
			if (x < 1) requestAnimationFrame(step)
			else if (gap) {
				setArc(gap, 0.99, 0.01)
				ring.classList.add('is-done')
			}
		}
		requestAnimationFrame(step)
	}

	// Hero chat: someone types a request, the assistant answers and proves the properties
	function chatPlayer(root) {
		const USER =
			'Draft a second theorem that states that successful signing payloads for two different signing domains cannot have equal bytes. ' +
			'Use the existing Lean model and the selected signing-domain theorem. Give the draft a new name. ' +
			'Include business conditions and an explanation for every complete input and the complete outcome. Do not attempt a proof.'
		const LINK =
			'<span class="cx-ref"><svg viewBox="0 0 16 16" aria-hidden="true"><path d="M8 1.8l5 1.9v4c0 3-2.1 5.4-5 6.5-2.9-1.1-5-3.5-5-6.5v-4z M5.6 8l1.7 1.7 3.2-3.4" /></svg>Signing domain separation <span class="cx-ref__arrow">&#8599;</span></span>'
		// reply tokens: plain words, plus the theorem reference as one piece
		const REPLY = ['This', 'proposal', 'restates', LINK].concat(
			'with explicit successful payload results. The domains must differ, and both payload calls must succeed. The resulting byte sequences must differ. This draft has not been checked by Lean.'.split(' ')
		)
		const typed = root.querySelector('.cx-composer__typed')
		const send = root.querySelector('.cx-composer__send')
		const bubble = root.querySelector('.cx-bubble')
		const reply = root.querySelector('.cx-reply')
		const replyText = root.querySelector('.cx-reply__text')
		const toggle = root.querySelector('.cx-chat__toggle')
		let paused = false
		let run = 0

		// a delay that stops counting while paused, and is abandoned on replay
		const wait = (ms, id) =>
			new Promise((resolve) => {
				let left = ms
				let last = performance.now()
				const tick = () => {
					if (id !== run) return
					const now = performance.now()
					if (!paused) left -= now - last
					last = now
					if (left <= 0) resolve()
					else setTimeout(tick, Math.min(16, Math.max(0, left)))
				}
				setTimeout(tick, Math.min(16, ms))
			})

		const reset = () => {
			root.classList.remove('is-typing', 'has-text', 'is-ready', 'is-done', 'is-paused', 'has-msgs')
			paused = false
			typed.textContent = ''
			bubble.hidden = true
			reply.hidden = true
			replyText.textContent = ''
		}

		const finalState = () => {
			reset()
			bubble.textContent = USER
			bubble.hidden = false
			root.classList.add('has-msgs')
			reply.hidden = false
			replyText.innerHTML = REPLY.join(' ')
			root.classList.add('is-done')
		}

		async function play() {
			const id = ++run
			reset()
			await wait(500, id)
			root.classList.add('is-typing')
			// fast typing: a few characters at a time
			for (let i = 0; i < USER.length; ) {
				i = Math.min(USER.length, i + 2 + Math.floor(Math.random() * 3))
				typed.textContent = USER.slice(0, i)
				root.classList.add('has-text')
				await wait(18, id)
			}
			// typing finished: the send button becomes active, then is pressed
			root.classList.remove('is-typing')
			root.classList.add('is-ready')
			await wait(550, id)
			send.classList.add('is-pressed')
			await wait(140, id)
			send.classList.remove('is-pressed')
			root.classList.remove('is-typing', 'has-text', 'is-ready')
			typed.textContent = ''
			bubble.textContent = USER
			bubble.hidden = false
			root.classList.add('has-msgs')
			await wait(450, id)
			reply.hidden = false
			replyText.innerHTML = '<span class="cx-dots"><span></span><span></span><span></span></span>'
			await wait(1100, id)
			for (let i = 1; i <= REPLY.length; i++) {
				replyText.innerHTML = REPLY.slice(0, i).join(' ')
				await wait(40, id)
			}
			root.classList.add('is-done')
		}

		toggle.addEventListener('click', () => {
			if (root.classList.contains('is-done')) return play()
			paused = !paused
			root.classList.toggle('is-paused', paused)
			toggle.setAttribute('aria-label', paused ? 'Play' : 'Pause')
		})

		return { play, finalState }
	}

	function init() {
		document.querySelectorAll('.cx-ring__fill, .cx-ring__gap').forEach((el) => setArc(el, 0, 0))

		const io = new IntersectionObserver(
			(entries) =>
				entries.forEach((e) => {
					if (!e.isIntersecting) return
					e.target.classList.add('is-in')
					if (e.target.dataset.ring && !e.target.dataset.played) {
						e.target.dataset.played = '1'
						setTimeout(() => playRing(e.target), reduced ? 0 : 300)
					}
					io.unobserve(e.target)
				}),
			{ threshold: 0.3 }
		)
		document.querySelectorAll('.cx-reveal').forEach((el) => io.observe(el))

		// start the chat once it is on screen
		const chatEl = document.querySelector('[data-chat]')
		if (chatEl) {
			const chat = chatPlayer(chatEl)
			if (reduced) chat.finalState()
			else {
				const co = new IntersectionObserver((es) => {
					if (es.some((e) => e.isIntersecting)) {
						co.disconnect()
						chat.play()
					}
				}, { threshold: 0.4 })
				co.observe(chatEl)
			}
		}

		// How it works: the line fills with the scroll and lights the three steps in turn
		const steps = document.querySelector('[data-scene="steps"]')
		if (!steps) return
		const nodes = steps.querySelectorAll('.cx-node')
		const copies = steps.querySelectorAll('.cx-steps__copy p')
		const fill = steps.querySelector('.cx-track__fill')
		const clamp = (x) => Math.max(0, Math.min(1, x))
		const seg = (p, a, b) => clamp((p - a) / (b - a))
		const at = [0.08, 0.4, 0.72]
		let ticking = false
		const render = () => {
			ticking = false
			const r = steps.getBoundingClientRect()
			const p = reduced ? 1 : clamp(-r.top / (r.height - window.innerHeight))
			fill.style.transform = `scaleX(${ease(seg(p, 0.08, 0.72))})`
			nodes.forEach((n, i) => n.classList.toggle('is-on', p >= at[i]))
			copies.forEach((c, i) => {
				const inn = ease(seg(p, at[i], at[i] + 0.06))
				const out = i < 2 ? 1 - ease(seg(p, at[i + 1] - 0.05, at[i + 1])) : 1
				const o = Math.min(inn, out)
				c.style.opacity = o
				c.style.transform = `translateY(${(1 - o) * 14}px)`
			})
		}
		const request = () => {
			if (!ticking) {
				ticking = true
				requestAnimationFrame(render)
			}
		}
		window.addEventListener('scroll', request, { passive: true })
		window.addEventListener('resize', request)
		render()
	}

	if (document.readyState === 'loading') document.addEventListener('DOMContentLoaded', init)
	else init()
})()
