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
		const USER = 'Build a payment service where a refund can never exceed the original charge.'
		const REPLY = 'Done. I wrote the payment service and proved its properties in Lean.'
		const typed = root.querySelector('.cx-composer__typed')
		const send = root.querySelector('.cx-composer__send')
		const bubble = root.querySelector('.cx-bubble')
		const reply = root.querySelector('.cx-reply')
		const replyText = root.querySelector('.cx-reply__text')
		const props = root.querySelector('.cx-props')
		const rows = [...root.querySelectorAll('.cx-props__row')]
		const allProved = root.querySelector('.cx-props__head .cx-pill')
		const toggle = root.querySelector('.cx-chat__toggle')
		let paused = false
		let run = 0

		// a delay that stops counting while paused, and is abandoned on replay
		const wait = (ms, id) =>
			new Promise((resolve) => {
				let left = ms
				let last = performance.now()
				const tick = (now) => {
					if (id !== run) return
					if (!paused) left -= now - last
					last = now
					if (left <= 0) resolve()
					else requestAnimationFrame(tick)
				}
				requestAnimationFrame(tick)
			})

		const reset = () => {
			root.classList.remove('is-typing', 'has-text', 'is-done', 'is-paused', 'has-msgs')
			paused = false
			typed.textContent = ''
			bubble.hidden = true
			reply.hidden = true
			replyText.textContent = ''
			props.hidden = true
			rows.forEach((r) => r.classList.remove('is-on'))
			allProved.classList.remove('is-on')
		}

		const finalState = () => {
			reset()
			bubble.textContent = USER
			bubble.hidden = false
			root.classList.add('has-msgs')
			reply.hidden = false
			replyText.textContent = REPLY
			props.hidden = false
			rows.forEach((r) => r.classList.add('is-on'))
			allProved.classList.add('is-on')
			root.classList.add('is-done')
		}

		async function play() {
			const id = ++run
			reset()
			await wait(500, id)
			root.classList.add('is-typing')
			for (const ch of USER) {
				typed.textContent += ch
				root.classList.add('has-text')
				await wait(ch === ' ' ? 45 : 26 + Math.random() * 30, id)
			}
			await wait(450, id)
			send.classList.add('is-pressed')
			await wait(140, id)
			send.classList.remove('is-pressed')
			root.classList.remove('is-typing', 'has-text')
			typed.textContent = ''
			bubble.textContent = USER
			bubble.hidden = false
			root.classList.add('has-msgs')
			await wait(450, id)
			reply.hidden = false
			replyText.innerHTML = '<span class="cx-dots"><span></span><span></span><span></span></span>'
			await wait(1100, id)
			replyText.textContent = ''
			const words = REPLY.split(' ')
			for (let i = 0; i < words.length; i++) {
				replyText.textContent += (i ? ' ' : '') + words[i]
				await wait(55, id)
			}
			await wait(300, id)
			props.hidden = false
			for (const r of rows) {
				await wait(380, id)
				r.classList.add('is-on')
			}
			await wait(350, id)
			allProved.classList.add('is-on')
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
