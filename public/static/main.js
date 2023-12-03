document.addEventListener("DOMContentLoaded", () => {
	const imagesLoading = document.querySelectorAll('[data-loading]');
	imagesLoading.forEach(img => {
		const loadingContainer = document.createElement('div');
		loadingContainer.setAttribute('class', 'loading-container');
		const skeleton = document.createElement('div');
		skeleton.setAttribute('class', 'skeleton');
		skeleton.setAttribute('style', 'width: ' + img.width + 'px; height: ' + img.height + 'px;');
		loadingContainer.appendChild(skeleton);

		img.insertAdjacentElement('beforebegin', loadingContainer);

		img.onload = () => {
			img.removeAttribute('data-loading');
			if (loadingContainer) {
				loadingContainer.remove();
			}
		};

		// Check if the image is already cached
		if (img.complete && img.naturalWidth !== 0) {
			img.onload();
		}
	});
});
