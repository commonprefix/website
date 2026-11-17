// Scroll-driven scenes for the Common Prefix Coder page.
;(function () {
	const clamp = (x, a = 0, b = 1) => Math.max(a, Math.min(b, x))
	const seg = (p, a, b) => clamp((p - a) / (b - a))
	const ease = (x) => (x < 0.5 ? 4 * x * x * x : 1 - Math.pow(-2 * x + 2, 3) / 2)
	const lerp = (a, b, x) => a + (b - a) * x
	const fade = (p, a, b, c, d) => Math.min(ease(seg(p, a, b)), c == null ? 1 : 1 - ease(seg(p, c, d)))
	const reduced = window.matchMedia('(prefers-reduced-motion: reduce)').matches

	function init() {
		const root = document.documentElement
		const logo = document.querySelector('.header__logo')
		if (logo) logo.src = '/static/coder/cp-text-logo-light.svg'


		// Reveal-on-scroll elements
		const io = new IntersectionObserver(
			(es) => es.forEach((e) => e.isIntersecting && e.target.classList.add('is-in')),
			{ threshold: 0.25 }
		)
		document.querySelectorAll('.cx-reveal').forEach((el) => io.observe(el))

		// Statement: one span per word
		const st = document.querySelector('.cx-statement')
		st.innerHTML = st.textContent.trim().split(' ').map((w) => `<span>${w}</span>`).join(' ')
		const words = [...st.children]

		// Ring scene
		const C = 2 * Math.PI * 300
		const arc = (el, from, len) => {
			el.style.strokeDasharray = `${Math.max(0, len) * C} ${C * 2}`
			el.style.strokeDashoffset = `${-from * C}`
			el.style.opacity = len <= 0 ? 0 : 1
		}
		const ring = document.querySelector('[data-scene="ring"]')
		const R = {
			soft: ring.querySelector('.cx-ring__soft'),
			cams: ring.querySelectorAll('.cx-cam'),
			dark: ring.querySelector('.cx-ring__dark'),
			green: ring.querySelector('.cx-ring__green'),
			gap: ring.querySelector('.cx-ring__gap'),
			count: ring.querySelector('.cx-ring__count'),
			arrow: ring.querySelector('.cx-ring__arrow'),
			titles: ring.querySelectorAll('.cx-ring__title'),
			subs: ring.querySelectorAll('.cx-ring__sub'),
		}
		const ang = ((-90 + 360 * 0.995) * Math.PI) / 180
		const P = [500 + 300 * Math.cos(ang), 500 + 300 * Math.sin(ang)]

		// Steps scene
		const steps = document.querySelector('[data-scene="steps"]')
		const nodes = steps.querySelectorAll('.cx-node')
		const copies = steps.querySelectorAll('.cx-steps__copy p')
		const fill = steps.querySelector('.cx-track__fill')

		const progress = (el) => {
			const r = el.getBoundingClientRect()
			return clamp(-r.top / (r.height - window.innerHeight))
		}

		function renderStatement(p) {
			words.forEach((w, i) => {
				const x = seg(p, 0.1 + (i / words.length) * 0.65, 0.1 + ((i + 1) / words.length) * 0.65)
				w.style.opacity = lerp(0.15, 1, x)
			})
		}

		function renderRing(p) {
			// 0-.22 fill to 99%, .22-.27 red 1%, .27-.38 zoom in, .38-.72 two lines, .72-.8 zoom out, .8-1 green to 100%
			const fillP = 0.99 * ease(seg(p, 0.02, 0.22))
			const red = ease(seg(p, 0.22, 0.26))
			const z = ease(seg(p, 0.27, 0.38)) * (1 - ease(seg(p, 0.72, 0.8)))
			const toDark = ease(seg(p, 0.78, 0.82))
			const sweep = ease(seg(p, 0.83, 0.97))

			arc(R.dark, 0, toDark >= 1 ? 1 : fillP)
			arc(R.gap, 0.99, 0.01 * red)
			R.gap.style.stroke = toDark > 0 ? `color-mix(in srgb, #c0322b ${100 - toDark * 100}%, #2a221d)` : ''
			R.gap.style.opacity = toDark >= 1 ? 0 : red > 0 ? 1 : 0
			arc(R.green, 0, sweep)
			R.count.textContent =
				p < 0.8 ? Math.round(fillP * 100) + '%' : (sweep <= 0 ? 0 : Math.max(1, Math.round(sweep * 100))) + '%'
			R.count.style.opacity = p < 0.8 ? 1 - seg(p, 0.27, 0.3) : seg(p, 0.8, 0.83)
			R.count.style.fill = sweep > 0 ? 'var(--cx-green)' : ''

			const k = lerp(1, 4, z)
			const Q = [lerp(P[0], 500, z), lerp(P[1], 560, z)]
			const tr = `translate(${Q[0] - k * P[0]} ${Q[1] - k * P[1]}) scale(${k})`
			R.cams.forEach((g) => g.setAttribute('transform', tr))
			R.soft.style.filter = `blur(${z * 7}px)`
			R.soft.style.opacity = lerp(1, 0.5, z)

			const a = ease(seg(p, 0.38, 0.43)) * (1 - ease(seg(p, 0.72, 0.75)))
			R.arrow.style.opacity = a
			R.arrow.style.transform = `translateY(${(1 - a) * -16}px)`

			const show = (el, o) => {
				el.style.opacity = o
				el.style.transform = `translateY(${(1 - o) * 12}px)`
			}
			show(R.titles[0], fade(p, 0, 0.03, 0.2, 0.23))
			show(R.titles[1], fade(p, 0.22, 0.26, 0.72, 0.75))
			show(R.titles[2], fade(p, 0.8, 0.84))
			show(R.subs[0], fade(p, 0.38, 0.43, 0.53, 0.56))
			show(R.subs[1], fade(p, 0.56, 0.6, 0.72, 0.75))
			show(R.subs[2], fade(p, 0.9, 0.95))
		}

		function renderSteps(p) {
			const at = [0.12, 0.42, 0.72]
			fill.style.transform = `scaleX(${ease(seg(p, 0.08, 0.78))})`
			nodes.forEach((n, i) => n.classList.toggle('is-on', p >= at[i]))
			copies.forEach((c, i) => {
				const o = fade(p, at[i], at[i] + 0.06, i < 2 ? at[i + 1] - 0.04 : 2, i < 2 ? at[i + 1] : 3)
				c.style.opacity = o
				c.style.transform = `translateY(${(1 - o) * 14}px)`
			})
		}

		function frame() {
			ticking = false
			renderStatement(reduced ? 1 : progress(st.closest('.cx-scene')))
			renderRing(reduced ? 1 : progress(ring))
			renderSteps(reduced ? 1 : progress(steps))
		}
		let ticking = false
		const request = () => {
			if (!ticking) {
				ticking = true
				requestAnimationFrame(frame)
			}
		}
		window.addEventListener('scroll', request, { passive: true })
		window.addEventListener('resize', request)
		frame()
		document.documentElement.classList.add('cx-ready')
	}

	if (document.readyState === 'loading') document.addEventListener('DOMContentLoaded', init)
	else init()
})()
