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
