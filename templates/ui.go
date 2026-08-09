package templates

import (
	"encoding/json"
	"fmt"
	"nyumba/models" // Standardizing to the central model package
)



func getHeader() string {
	return `
	<header class="fixed top-0 left-0 w-full z-50 backdrop-blur-xl bg-black/40 border-b border-white/5">
		<nav class="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
			<div class="flex items-center gap-12">
				<a href="/" class="text-2xl font-black tracking-tighter">Nyumba.</a>
				<div class="hidden md:flex items-center gap-8 text-sm font-bold text-zinc-400">
					<a href="/" class="hover:text-white transition-colors">Home</a>
					<a href="/explore" class="hover:text-white transition-colors">Listings</a>
					<a href="/about" class="hover:text-white transition-colors">About</a>
					<a href="/contact" class="hover:text-white transition-colors">Contact</a>
				</div>
			</div>
			<div class="flex items-center gap-4">
				<a href="/login" class="hidden md:block text-sm font-bold text-zinc-400 hover:text-white transition-colors">Login</a>
				<a href="/signup" class="bg-white text-black px-6 py-2.5 rounded-full text-sm font-bold hover:scale-105 transition-all">Sign Up</a>
				<button id="mobile-menu-btn" class="md:hidden p-2 text-zinc-400">
					<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
				</button>
			</div>
		</nav>
		<div id="mobile-menu" class="hidden md:hidden bg-zinc-950 border-b border-white/5 p-6 flex flex-col gap-4">
			<a href="/" class="text-lg font-bold">Home</a>
			<a href="/explore" class="text-lg font-bold">Listings</a>
			<a href="/about" class="text-lg font-bold">About</a>
			<a href="/contact" class="text-lg font-bold">Contact</a>
			<hr class="border-white/5">
			<a href="/login" class="text-lg font-bold">Login</a>
		</div>
	</header>
	<script>
		document.getElementById('mobile-menu-btn').addEventListener('click', () => {
			document.getElementById('mobile-menu').classList.toggle('hidden');
		});
	</script>`
}

func getFooter() string {
	return `
	<footer class="bg-zinc-950 border-t border-white/5 py-20">
		<div class="max-w-7xl mx-auto px-6 grid grid-cols-1 md:grid-cols-4 gap-12">
			<div class="col-span-1 md:col-span-2">
				<h2 class="text-3xl font-black tracking-tighter mb-6">Nyumba.</h2>
				<p class="text-zinc-500 max-w-sm leading-relaxed mb-8">Kenya's premier sanctuary discovery platform. Connecting serious renters with verified landlords directly.</p>
				<div class="flex gap-4">
					<a href="#" class="p-3 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/></svg></a>
					<a href="#" class="p-3 bg-white/5 rounded-2xl hover:bg-white/10 transition-colors"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="20" height="20" x="2" y="2" rx="5" ry="5"/><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"/><line x1="17.5" x2="17.51" y1="6.5" y2="6.5"/></svg></a>
				</div>
			</div>
			<div>
				<h4 class="text-sm font-black uppercase tracking-widest text-zinc-500 mb-6">Platform</h4>
				<ul class="space-y-4 text-zinc-400 font-bold">
					<li><a href="/explore" class="hover:text-white transition-colors">Listings</a></li>
					<li><a href="/landlord" class="hover:text-white transition-colors">For Landlords</a></li>
					<li><a href="/about" class="hover:text-white transition-colors">How it Works</a></li>
				</ul>
			</div>
			<div>
				<h4 class="text-sm font-black uppercase tracking-widest text-zinc-500 mb-6">Support</h4>
				<ul class="space-y-4 text-zinc-400 font-bold">
					<li><a href="/contact" class="hover:text-white transition-colors">Contact Us</a></li>
					<li><a href="#" class="hover:text-white transition-colors">Privacy Policy</a></li>
					<li><a href="#" class="hover:text-white transition-colors">Terms of Service</a></li>
				</ul>
			</div>
		</div>
		<div class="max-w-7xl mx-auto px-6 mt-20 pt-8 border-t border-white/5 text-center text-zinc-600 text-sm font-bold">
			&copy; 2026 Nyumba Technologies. All rights reserved.
		</div>
	</footer>`
}

func GetLandingHTML(featured []models.House) string {
	featuredJSON, _ := json.Marshal(featured)
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Nyumba - Find Your Sanctuary</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
	<style>
		body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
		@keyframes marquee { 0%% { transform: translateX(0); } 100%% { transform: translateX(-50%%); } }
		.marquee-container { overflow: hidden; white-space: nowrap; position: relative; }
		.marquee-content { display: inline-block; animation: marquee 40s linear infinite; }
		.bg-mesh {
			background-image: 
				radial-gradient(at 0%% 0%%, rgba(30, 58, 138, 0.3) 0px, transparent 50%%),
				radial-gradient(at 100%% 0%%, rgba(20, 184, 166, 0.2) 0px, transparent 50%%);
		}
	</style>
</head>
<body class="bg-mesh">
	%s

	<section class="pt-48 pb-32 flex flex-col items-center text-center px-6">
		<div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-white/10 bg-white/5 text-[10px] font-black tracking-widest uppercase mb-12">
			<span class="w-2 h-2 rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.6)]"></span>
			Verified Listings Only
		</div>
		<h1 class="text-6xl md:text-9xl font-black tracking-tighter mb-8 leading-[0.85]">
			Find Your <br> <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-cyan-400 to-emerald-400">Sanctuary.</span>
		</h1>
		<p class="max-w-xl text-lg md:text-xl text-zinc-400 mb-12 leading-relaxed font-medium">
			An exclusive platform connecting serious renters with verified landlords. No agents. No endless scrolling. Just your next home.
		</p>
		<div class="flex flex-col md:flex-row gap-4">
			<a href="/explore" class="group bg-white text-black px-10 py-5 rounded-full text-lg font-black hover:scale-105 transition-all flex items-center gap-3 shadow-[0_20px_50px_rgba(255,255,255,0.1)]">
				Start Your Search
				<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" class="group-hover:translate-x-1 transition-transform"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
			</a>
			<a href="/landlord" class="px-10 py-5 rounded-full text-lg font-black border border-white/10 hover:bg-white/5 transition-all">List Your Property</a>
		</div>
	</section>

	<section class="py-32 bg-white/5 border-y border-white/5">
		<div class="max-w-7xl mx-auto px-6">
			<div class="flex justify-between items-end mb-16">
				<div>
					<h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black mb-4">Featured Sanctuaries</h3>
					<h2 class="text-5xl font-black tracking-tighter">Hand-picked for you</h2>
				</div>
				<a href="/explore" class="text-zinc-400 hover:text-white font-bold underline">View all listings</a>
			</div>
			<div id="featured-grid" class="grid grid-cols-1 md:grid-cols-3 gap-8"></div>
		</div>
	</section>

	%s

	<script>
		const featured = %s;
		const grid = document.getElementById('featured-grid');
		featured.forEach(house => {
			const card = document.createElement('div');
			card.className = 'bg-white/5 backdrop-blur-xl rounded-[2.5rem] overflow-hidden border border-white/10 hover:border-white/20 transition-all group';
			card.innerHTML = %s
				<div class="aspect-video overflow-hidden relative">
					<img src="${house.image_urls[0]}" alt="${house.building_name}" class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700">
					<div class="absolute top-6 left-6 bg-blue-600 text-white px-4 py-1.5 rounded-full text-[10px] font-black tracking-widest uppercase">
						${house.location}
					</div>
				</div>
				<div class="p-8">
					<h3 class="text-2xl font-black tracking-tighter mb-2">${house.building_name}</h3>
					<div class="flex justify-between items-center pt-6 border-t border-white/5">
						<span class="text-xl font-black">KSh ${house.price.toLocaleString()}</span>
						<a href="/explore" class="text-sm font-black text-blue-400 hover:text-blue-300 transition-colors">View Details</a>
					</div>
				</div>
			%s;
			grid.appendChild(card);
		});
	</script>
</body>
</html>`, getHeader(), getFooter(), featuredJSON, "`", "`")
}

func GetExploreHTML(houses []models.House) string {
	initialJSON, _ := json.Marshal(houses)
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Explore Sanctuaries - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
	<style>
		body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
	</style>
</head>
<body class="min-h-screen flex flex-col justify-between">
	%s
	<main class="max-w-7xl mx-auto px-6 pt-32 pb-20 w-full">
		<div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-12">
			<div>
				<h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black mb-3">Curated Selection</h3>
				<h1 class="text-5xl md:text-7xl font-black tracking-tighter leading-none">Available Sanctuaries</h1>
			</div>
			<div id="results-count" class="text-zinc-400 text-sm font-bold">Showing verified properties</div>
		</div>

		<!-- Search & Filter Controls -->
		<div class="bg-white/5 backdrop-blur-xl p-6 rounded-3xl border border-white/10 mb-12 shadow-2xl">
			<form id="filter-form" class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-4">
				<div>
					<label class="block text-[10px] uppercase font-black tracking-widest text-zinc-400 mb-2">County (47 Counties)</label>
					<select id="filter-county" class="w-full bg-black/40 border border-white/10 rounded-2xl px-4 py-3 text-sm font-semibold focus:outline-none focus:border-blue-500 transition-colors text-white">
						<option value="">All Kenya Counties</option>
					</select>
				</div>
				<div>
					<label class="block text-[10px] uppercase font-black tracking-widest text-zinc-400 mb-2">Sub-County</label>
					<select id="filter-subcounty" class="w-full bg-black/40 border border-white/10 rounded-2xl px-4 py-3 text-sm font-semibold focus:outline-none focus:border-blue-500 transition-colors text-white">
						<option value="">All Sub-Counties</option>
					</select>
				</div>
				<div>
					<label class="block text-[10px] uppercase font-black tracking-widest text-zinc-400 mb-2">Town / City</label>
					<select id="filter-town" class="w-full bg-black/40 border border-white/10 rounded-2xl px-4 py-3 text-sm font-semibold focus:outline-none focus:border-blue-500 transition-colors text-white">
						<option value="">All Towns</option>
					</select>
				</div>
				<div>
					<label class="block text-[10px] uppercase font-black tracking-widest text-zinc-400 mb-2">Keyword Search</label>
					<input type="text" id="filter-location" placeholder="e.g. Westlands, Kilimani..." class="w-full bg-black/40 border border-white/10 rounded-2xl px-4 py-3 text-sm font-semibold focus:outline-none focus:border-blue-500 transition-colors">
				</div>
				<div>
					<label class="block text-[10px] uppercase font-black tracking-widest text-zinc-400 mb-2">Max Rent (KES)</label>
					<input type="number" id="filter-maxprice" placeholder="e.g. 150000" class="w-full bg-black/40 border border-white/10 rounded-2xl px-4 py-3 text-sm font-semibold focus:outline-none focus:border-blue-500 transition-colors">
				</div>
				<div class="flex items-end gap-2">
					<button type="submit" class="flex-1 bg-white text-black font-black py-3 rounded-2xl text-sm hover:scale-105 transition-all">Search</button>
					<button type="button" id="reset-filters" class="bg-white/10 text-zinc-300 font-bold px-4 py-3 rounded-2xl text-sm hover:bg-white/20 transition-all">Reset</button>
				</div>
			</form>
		</div>

		<!-- Grid States -->
		<div id="loading-state" class="hidden py-24 text-center">
			<div class="inline-block w-12 h-12 border-4 border-white/20 border-t-white rounded-full animate-spin mb-4"></div>
			<p class="text-zinc-400 font-bold text-sm">Searching available sanctuaries...</p>
		</div>

		<div id="error-state" class="hidden py-16 bg-red-500/10 border border-red-500/20 rounded-3xl p-8 text-center">
			<p class="text-red-400 font-bold text-lg mb-4" id="error-message">Failed to load properties.</p>
			<button id="retry-btn" class="bg-red-500 text-white font-bold px-6 py-2.5 rounded-full text-sm hover:bg-red-600 transition-colors">Retry Connection</button>
		</div>

		<div id="empty-state" class="hidden py-24 text-center border border-dashed border-white/10 rounded-3xl p-12">
			<h3 class="text-2xl font-black mb-2">No matching sanctuaries found</h3>
			<p class="text-zinc-400 text-sm mb-6 max-w-md mx-auto font-medium">Try broadening your location search or raising the max rent filter.</p>
			<button onclick="document.getElementById('reset-filters').click()" class="bg-white text-black px-6 py-3 rounded-full text-sm font-black hover:scale-105 transition-all">Clear Filters</button>
		</div>

		<div id="houses-grid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"></div>
	</main>

	<!-- Action Modals (Inquire / Book / Payment) -->
	<div id="action-modal" class="hidden fixed inset-0 z-50 bg-black/80 backdrop-blur-md flex items-center justify-center p-6">
		<div class="bg-zinc-900 border border-white/10 rounded-3xl p-8 max-w-lg w-full relative">
			<button id="close-modal-btn" class="absolute top-6 right-6 p-2 text-zinc-400 hover:text-white">&times;</button>
			<div id="modal-content"></div>
		</div>
	</div>

	%s

	<script>
		let currentHouses = %s;

		function renderHouses(houses) {
			const grid = document.getElementById('houses-grid');
			const emptyState = document.getElementById('empty-state');
			const countLabel = document.getElementById('results-count');

			grid.innerHTML = '';
			if (!houses || houses.length === 0) {
				emptyState.classList.remove('hidden');
				grid.classList.add('hidden');
				countLabel.textContent = '0 properties found';
				return;
			}

			emptyState.classList.add('hidden');
			grid.classList.remove('hidden');
			countLabel.textContent = `Showing ${houses.length} properties`;

			houses.forEach(house => {
				const card = document.createElement('div');
				card.className = 'bg-white/5 backdrop-blur-xl rounded-[2.5rem] overflow-hidden border border-white/10 hover:border-white/20 transition-all flex flex-col justify-between group';
				
				const imgUrl = (house.image_urls && house.image_urls.length > 0) ? house.image_urls[0] : 'https://images.unsplash.com/photo-1600585154340-be6161a56a0c?auto=format&fit=crop&w=1200&q=80';
				
				card.innerHTML = `
					<div>
						<div class="aspect-video overflow-hidden relative">
							<img src="${imgUrl}" alt="${house.building_name}" class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500">
							<div class="absolute top-4 left-4 bg-black/60 backdrop-blur-md border border-white/10 text-white px-3 py-1 rounded-full text-[10px] font-black tracking-widest uppercase">
								${house.location}
							</div>
							${house.is_paid ? '<div class="absolute top-4 right-4 bg-emerald-500 text-black font-black px-3 py-1 rounded-full text-[10px] tracking-widest uppercase">VERIFIED</div>' : ''}
						</div>
						<div class="p-6">
							<h3 class="text-2xl font-black tracking-tighter mb-2">${house.building_name}</h3>
							<p class="text-zinc-400 text-xs font-medium line-clamp-2 mb-4">${house.description || 'Modern sanctuary with luxury finishes.'}</p>
							<div class="flex gap-4 text-xs font-bold text-zinc-400 mb-6">
								<span>🛏️ ${house.bedrooms} Bed</span>
								<span>🚿 ${house.bathrooms} Bath</span>
							</div>
						</div>
					</div>
					<div class="p-6 pt-0 border-t border-white/5 flex flex-col gap-3">
						<div class="flex justify-between items-center mt-4">
							<span class="text-xl font-black">KES ${Number(house.price).toLocaleString()}</span>
							<button onclick="toggleFavorite(${house.id})" class="p-2.5 rounded-full bg-white/5 hover:bg-white/10 transition-colors text-xs font-bold">❤️ Favorite</button>
						</div>
						<div class="grid grid-cols-2 gap-2 mt-2">
							<button onclick="openInquireModal(${house.id}, '${house.building_name}')" class="bg-white/10 text-white hover:bg-white/20 py-2.5 rounded-xl text-xs font-bold transition-all">Inquire</button>
							<button onclick="openBookingModal(${house.id}, '${house.building_name}')" class="bg-white/10 text-white hover:bg-white/20 py-2.5 rounded-xl text-xs font-bold transition-all">Book Tour</button>
						</div>
						<button onclick="openPaymentModal(${house.id}, ${house.price}, '${house.building_name}')" class="w-full bg-emerald-500/20 border border-emerald-500/30 text-emerald-400 hover:bg-emerald-500/30 font-black py-2.5 rounded-xl text-xs transition-all flex items-center justify-center gap-2">
							💳 Pay Rent via M-Pesa STK
						</button>
					</div>
				`;
				grid.appendChild(card);
			});
		}

		async function loadLocationsUI() {
			try {
				const res = await fetch('/api/locations/counties');
				if (res.ok) {
					const counties = await res.json();
					const countySelect = document.getElementById('filter-county');
					if (countySelect && counties) {
						countySelect.innerHTML = '<option value="">All Kenya Counties</option>';
						counties.forEach(c => {
							const opt = document.createElement('option');
							opt.value = c.id;
							opt.textContent = `${c.code.toString().padStart(3, '0')} - ${c.name}`;
							countySelect.appendChild(opt);
						});
					}
				}
			} catch (err) {
				console.log('Location API notice:', err);
			}
		}

		document.getElementById('filter-county').addEventListener('change', async (e) => {
			const countyId = e.target.value;
			const subSelect = document.getElementById('filter-subcounty');
			const townSelect = document.getElementById('filter-town');

			subSelect.innerHTML = '<option value="">All Sub-Counties</option>';
			townSelect.innerHTML = '<option value="">All Towns</option>';

			if (!countyId) return;

			try {
				const [subRes, townRes] = await Promise.all([
					fetch('/api/locations/sub-counties?county_id=' + countyId),
					fetch('/api/locations/towns?county_id=' + countyId)
				]);
				if (subRes.ok) {
					const subs = await subRes.json();
					subs.forEach(s => {
						const opt = document.createElement('option');
						opt.value = s.id;
						opt.textContent = s.name;
						subSelect.appendChild(opt);
					});
				}
				if (townRes.ok) {
					const towns = await townRes.json();
					towns.forEach(t => {
						const opt = document.createElement('option');
						opt.value = t.id;
						opt.textContent = t.name;
						townSelect.appendChild(opt);
					});
				}
			} catch (err) {
				console.error('Failed to update sub-counties/towns', err);
			}
		});

		async function fetchHouses() {
			const loading = document.getElementById('loading-state');
			const error = document.getElementById('error-state');
			const grid = document.getElementById('houses-grid');

			loading.classList.remove('hidden');
			error.classList.add('hidden');
			grid.classList.add('hidden');

			const countyId = document.getElementById('filter-county').value;
			const subCountyId = document.getElementById('filter-subcounty').value;
			const townId = document.getElementById('filter-town').value;
			const loc = document.getElementById('filter-location').value.trim();
			const maxPrice = document.getElementById('filter-maxprice').value.trim();

			const params = new URLSearchParams();
			if (countyId) params.append('county_id', countyId);
			if (subCountyId) params.append('sub_county_id', subCountyId);
			if (townId) params.append('town_id', townId);
			if (loc) params.append('location', loc);
			if (maxPrice) params.append('max_price', maxPrice);

			try {
				const res = await fetch('/houses?' + params.toString());
				if (!res.ok) throw new Error(`HTTP error ${res.status}`);
				const data = await res.json();
				loading.classList.add('hidden');
				renderHouses(data);
			} catch (err) {
				loading.classList.add('hidden');
				error.classList.remove('hidden');
				document.getElementById('error-message').textContent = 'Unable to reach Nyumba backend servers.';
			}
		}

		document.getElementById('filter-form').addEventListener('submit', (e) => {
			e.preventDefault();
			fetchHouses();
		});

		document.getElementById('reset-filters').addEventListener('click', () => {
			document.getElementById('filter-county').value = '';
			document.getElementById('filter-subcounty').value = '<option value="">All Sub-Counties</option>';
			document.getElementById('filter-town').value = '<option value="">All Towns</option>';
			document.getElementById('filter-location').value = '';
			document.getElementById('filter-maxprice').value = '';
			fetchHouses();
		});

		document.getElementById('retry-btn').addEventListener('click', fetchHouses);

		// Initialize locations UI
		loadLocationsUI();

		// Modal Helpers
		const modal = document.getElementById('action-modal');
		const modalContent = document.getElementById('modal-content');
		document.getElementById('close-modal-btn').onclick = () => modal.classList.add('hidden');

		async function toggleFavorite(propertyId) {
			try {
				const res = await fetch('/favorites', {
					method: 'POST',
					headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
					body: new URLSearchParams({ property_id: propertyId })
				});
				if (res.status === 401) {
					alert('Please log in to save favorites.');
					window.location.href = '/login';
					return;
				}
				const data = await res.json();
				alert(data.message || 'Updated favorites!');
			} catch (err) {
				alert('Failed to update favorites');
			}
		}

		function openInquireModal(propertyId, name) {
			modalContent.innerHTML = `
				<h3 class="text-2xl font-black mb-2">Inquire about ${name}</h3>
				<p class="text-zinc-400 text-xs mb-6">Send a direct inquiry to the landlord.</p>
				<form id="inquiry-form" class="space-y-4">
					<textarea id="inquiry-msg" required rows="4" class="w-full bg-black/40 border border-white/10 rounded-2xl p-4 text-sm focus:outline-none focus:border-blue-500" placeholder="Hello, I am interested in viewing this property..."></textarea>
					<button type="submit" class="w-full bg-white text-black font-black py-3 rounded-2xl hover:scale-105 transition-all text-sm">Send Inquiry</button>
				</form>
			`;
			modal.classList.remove('hidden');

			document.getElementById('inquiry-form').onsubmit = async (e) => {
				e.preventDefault();
				const msg = document.getElementById('inquiry-msg').value;
				const res = await fetch('/inquiries', {
					method: 'POST',
					headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
					body: new URLSearchParams({ property_id: propertyId, message: msg })
				});
				if (res.status === 401) {
					alert('Please log in to send an inquiry.');
					window.location.href = '/login';
					return;
				}
				if (res.ok) {
					alert('Inquiry sent successfully!');
					modal.classList.add('hidden');
				} else {
					alert('Failed to send inquiry.');
				}
			};
		}

		function openBookingModal(propertyId, name) {
			modalContent.innerHTML = `
				<h3 class="text-2xl font-black mb-2">Book Inspection for ${name}</h3>
				<p class="text-zinc-400 text-xs mb-6">Schedule an in-person viewing with the landlord.</p>
				<form id="booking-form" class="space-y-4">
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-2">Preferred Date & Time</label>
						<input type="datetime-local" id="booking-date" required class="w-full bg-black/40 border border-white/10 rounded-2xl p-4 text-sm focus:outline-none focus:border-blue-500">
					</div>
					<button type="submit" class="w-full bg-white text-black font-black py-3 rounded-2xl hover:scale-105 transition-all text-sm">Confirm Booking</button>
				</form>
			`;
			modal.classList.remove('hidden');

			document.getElementById('booking-form').onsubmit = async (e) => {
				e.preventDefault();
				const dateVal = document.getElementById('booking-date').value;
				const res = await fetch('/bookings', {
					method: 'POST',
					headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
					body: new URLSearchParams({ property_id: propertyId, inspection_date: dateVal })
				});
				if (res.status === 401) {
					alert('Please log in to schedule a booking.');
					window.location.href = '/login';
					return;
				}
				if (res.ok) {
					alert('Inspection booking requested!');
					modal.classList.add('hidden');
				} else {
					alert('Failed to submit booking.');
				}
			};
		}

		function openPaymentModal(propertyId, amount, name) {
			modalContent.innerHTML = `
				<h3 class="text-2xl font-black mb-2">Pay via M-Pesa STK Push</h3>
				<p class="text-zinc-400 text-xs mb-6">Property: ${name} (KES ${Number(amount).toLocaleString()})</p>
				<form id="stk-form" class="space-y-4">
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-2">M-Pesa Phone Number</label>
						<input type="tel" id="stk-phone" placeholder="07XXXXXXXX" required class="w-full bg-black/40 border border-white/10 rounded-2xl p-4 text-sm focus:outline-none focus:border-emerald-500">
					</div>
					<div id="stk-status" class="hidden text-xs font-bold p-3 rounded-xl"></div>
					<button type="submit" id="stk-submit-btn" class="w-full bg-emerald-500 text-black font-black py-3 rounded-2xl hover:scale-105 transition-all text-sm">Initiate STK Push</button>
				</form>
			`;
			modal.classList.remove('hidden');

			document.getElementById('stk-form').onsubmit = async (e) => {
				e.preventDefault();
				const phone = document.getElementById('stk-phone').value;
				const statusDiv = document.getElementById('stk-status');
				const btn = document.getElementById('stk-submit-btn');

				statusDiv.classList.remove('hidden', 'bg-red-500/20', 'text-red-400', 'bg-emerald-500/20', 'text-emerald-400');
				statusDiv.classList.add('bg-blue-500/20', 'text-blue-400');
				statusDiv.textContent = 'Sending STK Push prompt to your phone...';
				btn.disabled = true;

				try {
					const res = await fetch('/api/payments/initiate', {
						method: 'POST',
						headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
						body: new URLSearchParams({ property_id: propertyId, amount: amount, phone: phone })
					});

					if (res.status === 401) {
						alert('Please log in to initiate payment.');
						window.location.href = '/login';
						return;
					}

					const data = await res.json();
					if (res.status === 202 && data.status === 'BLOCKED_EXTERNAL') {
						statusDiv.classList.replace('text-blue-400', 'text-yellow-400');
						statusDiv.textContent = 'Payment recorded locally in pending state (M-Pesa API credentials pending).';
					} else if (res.ok) {
						statusDiv.classList.replace('text-blue-400', 'text-emerald-400');
						statusDiv.textContent = 'STK Push sent! Please enter your PIN on your phone.';
					} else {
						statusDiv.classList.replace('text-blue-400', 'text-red-400');
						statusDiv.textContent = data.error || 'Failed to initiate payment.';
					}
				} catch (err) {
					statusDiv.classList.replace('text-blue-400', 'text-red-400');
					statusDiv.textContent = 'Network error initiating payment.';
				} finally {
					btn.disabled = false;
				}
			};
		}

		// Initial Render
		renderHouses(currentHouses);
	</script>
</body>
</html>`, getHeader(), getFooter(), initialJSON)
}

func GetLandlordHTML() string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Landlord Portal - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
	<style>
		body { font-family: 'Inter', sans-serif; background-color: #09090b; color: white; }
	</style>
</head>
<body class="min-h-screen flex flex-col justify-between">
	%s
	<main class="max-w-7xl mx-auto px-6 pt-32 pb-20 w-full">
		<div class="flex flex-col md:flex-row md:items-center justify-between gap-6 mb-12">
			<div>
				<h3 class="text-[10px] uppercase tracking-[0.3em] text-zinc-500 font-black mb-2">Management Console</h3>
				<h1 class="text-5xl md:text-7xl font-black tracking-tighter">Landlord Portal</h1>
			</div>
			<button id="add-property-btn" class="bg-white text-black px-8 py-4 rounded-full font-black hover:scale-105 transition-all text-sm flex items-center justify-center gap-2">
				+ Add New Property
			</button>
		</div>

		<!-- Dashboard Tabs -->
		<div class="flex border-b border-white/10 mb-8 gap-8">
			<button id="tab-props-btn" class="py-3 font-black text-sm text-white border-b-2 border-white">My Properties</button>
			<button id="tab-inq-btn" class="py-3 font-black text-sm text-zinc-400 hover:text-white transition-colors">Tenant Inquiries</button>
			<button id="tab-book-btn" class="py-3 font-black text-sm text-zinc-400 hover:text-white transition-colors">Scheduled Tours</button>
		</div>

		<!-- Properties Tab Content -->
		<div id="tab-props-content">
			<div id="landlord-props-loading" class="py-16 text-center">
				<div class="inline-block w-8 h-8 border-4 border-white/20 border-t-white rounded-full animate-spin"></div>
			</div>
			<div id="landlord-props-grid" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"></div>
		</div>

		<!-- Inquiries Tab Content -->
		<div id="tab-inq-content" class="hidden space-y-4">
			<div id="inquiries-list" class="space-y-4"></div>
		</div>

		<!-- Bookings Tab Content -->
		<div id="tab-book-content" class="hidden space-y-4">
			<div id="bookings-list" class="space-y-4"></div>
		</div>
	</main>

	<!-- Add Property Modal -->
	<div id="property-modal" class="hidden fixed inset-0 z-50 bg-black/80 backdrop-blur-md flex items-center justify-center p-6 overflow-y-auto">
		<div class="bg-zinc-900 border border-white/10 rounded-3xl p-8 max-w-xl w-full my-8 relative">
			<button id="close-prop-modal" class="absolute top-6 right-6 p-2 text-zinc-400 hover:text-white">&times;</button>
			<h3 class="text-2xl font-black mb-6" id="prop-modal-title">List New Property</h3>
			<form id="property-form" class="space-y-4">
				<input type="hidden" id="prop-id" name="property_id">
				<div>
					<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Building Name / Title</label>
					<input type="text" id="prop-building" required class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm focus:outline-none focus:border-blue-500">
				</div>
				<div class="grid grid-cols-2 gap-4">
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">County (Kenya 47)</label>
						<select id="prop-county" required class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm text-white focus:outline-none focus:border-blue-500">
							<option value="">Select County...</option>
						</select>
					</div>
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Sub-County</label>
						<select id="prop-subcounty" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm text-white focus:outline-none focus:border-blue-500">
							<option value="">Select Sub-County...</option>
						</select>
					</div>
				</div>
				<div class="grid grid-cols-2 gap-4">
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Town / City</label>
						<select id="prop-town" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm text-white focus:outline-none focus:border-blue-500">
							<option value="">Select Town...</option>
						</select>
					</div>
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Ward</label>
						<select id="prop-ward" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm text-white focus:outline-none focus:border-blue-500">
							<option value="">Select Ward...</option>
						</select>
					</div>
				</div>
				<div class="grid grid-cols-2 gap-4">
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Neighborhood / Street</label>
						<input type="text" id="prop-location" required placeholder="e.g. Westlands Rd, Kilimani" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm focus:outline-none focus:border-blue-500">
					</div>
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Monthly Price (KES)</label>
						<input type="number" id="prop-price" required class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm focus:outline-none focus:border-blue-500">
					</div>
				</div>
				<div class="grid grid-cols-2 gap-4">
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Bedrooms</label>
						<input type="number" id="prop-bedrooms" required min="1" value="1" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm focus:outline-none focus:border-blue-500">
					</div>
					<div>
						<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Bathrooms</label>
						<input type="number" id="prop-bathrooms" required min="1" value="1" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm focus:outline-none focus:border-blue-500">
					</div>
				</div>
				<div>
					<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Description</label>
					<textarea id="prop-desc" rows="3" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm focus:outline-none focus:border-blue-500"></textarea>
				</div>
				<div>
					<label class="block text-xs uppercase font-bold text-zinc-400 mb-1">Upload Photo</label>
					<input type="file" id="prop-image" accept="image/jpeg,image/png,image/webp" class="w-full bg-black/40 border border-white/10 rounded-xl p-3 text-sm text-zinc-400">
				</div>
				<button type="submit" class="w-full bg-white text-black font-black py-3.5 rounded-xl hover:scale-105 transition-all text-sm mt-4">Save Property</button>
			</form>
		</div>
	</div>

	%s

	<script>
		// Tab Switching Logic
		const tabPropsBtn = document.getElementById('tab-props-btn');
		const tabInqBtn = document.getElementById('tab-inq-btn');
		const tabBookBtn = document.getElementById('tab-book-btn');

		const contentProps = document.getElementById('tab-props-content');
		const contentInq = document.getElementById('tab-inq-content');
		const contentBook = document.getElementById('tab-book-content');

		function resetTabs() {
			[tabPropsBtn, tabInqBtn, tabBookBtn].forEach(b => {
				b.classList.remove('text-white', 'border-b-2', 'border-white');
				b.classList.add('text-zinc-400');
			});
			[contentProps, contentInq, contentBook].forEach(c => c.classList.add('hidden'));
		}

		tabPropsBtn.onclick = () => { resetTabs(); tabPropsBtn.classList.add('text-white', 'border-b-2', 'border-white'); contentProps.classList.remove('hidden'); loadLandlordProperties(); };
		tabInqBtn.onclick = () => { resetTabs(); tabInqBtn.classList.add('text-white', 'border-b-2', 'border-white'); contentInq.classList.remove('hidden'); loadInquiries(); };
		tabBookBtn.onclick = () => { resetTabs(); tabBookBtn.classList.add('text-white', 'border-b-2', 'border-white'); contentBook.classList.remove('hidden'); loadBookings(); };

		// Fetch Landlord's properties
		async function loadLandlordProperties() {
			const grid = document.getElementById('landlord-props-grid');
			const loader = document.getElementById('landlord-props-loading');
			loader.classList.remove('hidden');
			grid.innerHTML = '';

			try {
				const res = await fetch('/get-houses');
				if (res.status === 401) {
					window.location.href = '/login';
					return;
				}
				const houses = await res.json();
				loader.classList.add('hidden');

				if (!houses || houses.length === 0) {
					grid.innerHTML = '<p class="col-span-full text-center text-zinc-500 font-bold py-12">You have not listed any properties yet.</p>';
					return;
				}

				houses.forEach(house => {
					const card = document.createElement('div');
					card.className = 'bg-white/5 border border-white/10 rounded-3xl p-6 flex flex-col justify-between';
					const imgUrl = (house.image_urls && house.image_urls.length > 0) ? house.image_urls[0] : 'https://images.unsplash.com/photo-1600585154340-be6161a56a0c?auto=format&fit=crop&w=1200&q=80';

					card.innerHTML = `
						<div>
							<img src="${imgUrl}" class="aspect-video w-full object-cover rounded-2xl mb-4">
							<h3 class="text-xl font-black">${house.building_name}</h3>
							<p class="text-xs text-zinc-400 font-bold mb-4">${house.location} • KES ${Number(house.price).toLocaleString()}</p>
						</div>
						<div class="space-y-2 pt-4 border-t border-white/5">
							<div class="flex gap-2">
								${house.is_published 
									? `<button onclick="unpublishProp(${house.id})" class="flex-1 bg-yellow-500/20 text-yellow-400 py-2 rounded-xl text-xs font-bold">Unpublish</button>`
									: `<button onclick="publishProp(${house.id})" class="flex-1 bg-emerald-500/20 text-emerald-400 py-2 rounded-xl text-xs font-bold">Publish</button>`}
								<button onclick="deleteProp(${house.id})" class="bg-red-500/20 text-red-400 px-4 py-2 rounded-xl text-xs font-bold">Delete</button>
							</div>
						</div>
					`;
					grid.appendChild(card);
				});
			} catch (err) {
				loader.classList.add('hidden');
				grid.innerHTML = '<p class="col-span-full text-center text-red-400 font-bold">Failed to load property portfolio.</p>';
			}
		}

		async function publishProp(id) {
			await fetch('/publish-house', { method: 'POST', headers: { 'Content-Type': 'application/x-www-form-urlencoded' }, body: new URLSearchParams({ id }) });
			loadLandlordProperties();
		}

		async function unpublishProp(id) {
			await fetch('/unpublish-house', { method: 'POST', headers: { 'Content-Type': 'application/x-www-form-urlencoded' }, body: new URLSearchParams({ id }) });
			loadLandlordProperties();
		}

		async function deleteProp(id) {
			if (!confirm('Are you sure you want to delete this property listing?')) return;
			await fetch('/delete-house', { method: 'POST', headers: { 'Content-Type': 'application/x-www-form-urlencoded' }, body: new URLSearchParams({ id }) });
			loadLandlordProperties();
		}

		async function loadInquiries() {
			const container = document.getElementById('inquiries-list');
			container.innerHTML = '<p class="text-zinc-500 font-bold text-center py-8">Loading tenant inquiries...</p>';
			try {
				const res = await fetch('/inquiries');
				const data = await res.json();
				if (!data || data.length === 0) {
					container.innerHTML = '<p class="text-zinc-500 font-bold text-center py-8">No inquiries received yet.</p>';
					return;
				}
				container.innerHTML = data.map(inq => `
					<div class="bg-white/5 border border-white/10 p-6 rounded-2xl flex justify-between items-center">
						<div>
							<h4 class="font-black text-lg">Property #${inq.property_id}</h4>
							<p class="text-sm text-zinc-300 mt-1">${inq.message}</p>
							<span class="text-[10px] text-zinc-500 font-bold uppercase mt-2 inline-block">${new Date(inq.created_at).toLocaleDateString()}</span>
						</div>
						<span class="bg-blue-500/20 text-blue-400 px-3 py-1 rounded-full text-xs font-bold uppercase">${inq.status}</span>
					</div>
				`).join('');
			} catch (err) {
				container.innerHTML = '<p class="text-red-400 font-bold text-center py-8">Error loading inquiries.</p>';
			}
		}

		async function loadBookings() {
			const container = document.getElementById('bookings-list');
			container.innerHTML = '<p class="text-zinc-500 font-bold text-center py-8">Loading scheduled tours...</p>';
			try {
				const res = await fetch('/bookings');
				const data = await res.json();
				if (!data || data.length === 0) {
					container.innerHTML = '<p class="text-zinc-500 font-bold text-center py-8">No property viewing tours scheduled yet.</p>';
					return;
				}
				container.innerHTML = data.map(b => `
					<div class="bg-white/5 border border-white/10 p-6 rounded-2xl flex justify-between items-center">
						<div>
							<h4 class="font-black text-lg">Property #${b.property_id}</h4>
							<p class="text-xs text-zinc-400 font-bold mt-1">Inspection Date: ${new Date(b.inspection_date).toLocaleString()}</p>
						</div>
						<span class="bg-emerald-500/20 text-emerald-400 px-3 py-1 rounded-full text-xs font-bold uppercase">${b.status}</span>
					</div>
				`).join('');
			} catch (err) {
				container.innerHTML = '<p class="text-red-400 font-bold text-center py-8">Error loading scheduled tours.</p>';
			}
		}

		// Property Modal Handlers
		async function loadLandlordLocationsUI() {
			try {
				const res = await fetch('/api/locations/counties');
				if (res.ok) {
					const counties = await res.json();
					const countySelect = document.getElementById('prop-county');
					if (countySelect && counties) {
						countySelect.innerHTML = '<option value="">Select County...</option>';
						counties.forEach(c => {
							const opt = document.createElement('option');
							opt.value = c.id;
							opt.textContent = `${c.code.toString().padStart(3, '0')} - ${c.name}`;
							countySelect.appendChild(opt);
						});
					}
				}
			} catch (err) {
				console.log('Landlord location API notice:', err);
			}
		}

		document.getElementById('prop-county').addEventListener('change', async (e) => {
			const countyId = e.target.value;
			const subSelect = document.getElementById('prop-subcounty');
			const townSelect = document.getElementById('prop-town');
			const wardSelect = document.getElementById('prop-ward');

			subSelect.innerHTML = '<option value="">Select Sub-County...</option>';
			townSelect.innerHTML = '<option value="">Select Town...</option>';
			wardSelect.innerHTML = '<option value="">Select Ward...</option>';

			if (!countyId) return;

			try {
				const [subRes, townRes] = await Promise.all([
					fetch('/api/locations/sub-counties?county_id=' + countyId),
					fetch('/api/locations/towns?county_id=' + countyId)
				]);
				if (subRes.ok) {
					const subs = await subRes.json();
					subs.forEach(s => {
						const opt = document.createElement('option');
						opt.value = s.id;
						opt.textContent = s.name;
						subSelect.appendChild(opt);
					});
				}
				if (townRes.ok) {
					const towns = await townRes.json();
					towns.forEach(t => {
						const opt = document.createElement('option');
						opt.value = t.id;
						opt.textContent = t.name;
						townSelect.appendChild(opt);
					});
				}
			} catch (err) {
				console.error('Failed to update landlord sub-counties/towns', err);
			}
		});

		document.getElementById('prop-subcounty').addEventListener('change', async (e) => {
			const subId = e.target.value;
			const wardSelect = document.getElementById('prop-ward');
			wardSelect.innerHTML = '<option value="">Select Ward...</option>';

			if (!subId) return;

			try {
				const res = await fetch('/api/locations/wards?sub_county_id=' + subId);
				if (res.ok) {
					const wards = await res.json();
					wards.forEach(w => {
						const opt = document.createElement('option');
						opt.value = w.id;
						opt.textContent = w.name;
						wardSelect.appendChild(opt);
					});
				}
			} catch (err) {
				console.error('Failed to update wards', err);
			}
		});

		const propModal = document.getElementById('property-modal');
		document.getElementById('add-property-btn').onclick = () => {
			document.getElementById('property-form').reset();
			document.getElementById('prop-id').value = '';
			loadLandlordLocationsUI();
			propModal.classList.remove('hidden');
		};
		document.getElementById('close-prop-modal').onclick = () => propModal.classList.add('hidden');

		document.getElementById('property-form').onsubmit = async (e) => {
			e.preventDefault();
			const formData = new FormData();
			formData.append('building_name', document.getElementById('prop-building').value);
			formData.append('county_id', document.getElementById('prop-county').value);
			formData.append('sub_county_id', document.getElementById('prop-subcounty').value);
			formData.append('town_id', document.getElementById('prop-town').value);
			formData.append('ward_id', document.getElementById('prop-ward').value);
			formData.append('location', document.getElementById('prop-location').value);
			formData.append('price', document.getElementById('prop-price').value);
			formData.append('bedrooms', document.getElementById('prop-bedrooms').value);
			formData.append('bathrooms', document.getElementById('prop-bathrooms').value);
			formData.append('description', document.getElementById('prop-desc').value);

			const imgFile = document.getElementById('prop-image').files[0];
			if (imgFile) formData.append('property_photo', imgFile);

			const res = await fetch('/add-house', { method: 'POST', body: formData });
			if (res.ok) {
				alert('Property listed successfully!');
				propModal.classList.add('hidden');
				loadLandlordProperties();
			} else {
				const txt = await res.text();
				alert('Failed to save property: ' + txt);
			}
		};

		// Initial Load
		loadLandlordProperties();
	</script>
</body>
</html>`, getHeader(), getFooter())
}

func GetAuthHTML(mode string) string {
	isLogin := mode == "Login"
	
	// Determine form fields based on mode
	var formFields string
	if isLogin {
		formFields = `
			<div class="space-y-4">
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Email</label>
					<input type="email" name="email" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="your@email.com">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Password</label>
					<input type="password" name="password" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="••••••••">
				</div>
			</div>
			<button type="submit" class="w-full bg-white text-black font-black py-4 rounded-xl hover:scale-[1.02] transition-transform mt-6">Sign In</button>
			<p class="text-center text-zinc-500 text-sm mt-6 font-medium">Don't have an account? <a href="/signup" class="text-blue-400 hover:text-blue-300 font-bold">Sign up</a></p>
		`
	} else {
		formFields = `
			<div class="space-y-4">
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Full Name</label>
					<input type="text" name="name" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="Abdul Wandera">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Email</label>
					<input type="email" name="email" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="your@email.com">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Phone</label>
					<input type="tel" name="phone" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="+254 700 000 000">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Password</label>
					<input type="password" name="password" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:border-blue-500 transition-colors" placeholder="••••••••">
				</div>
				<div>
					<label class="block text-xs font-bold uppercase tracking-widest text-zinc-500 mb-2">Account Type</label>
					<select name="role" required class="w-full bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-blue-500 transition-colors">
						<option value="" disabled selected class="bg-zinc-900">Select account type</option>
						<option value="renter" class="bg-zinc-900">Looking for a home (Renter)</option>
						<option value="landlord" class="bg-zinc-900">Listing properties (Landlord)</option>
					</select>
				</div>
			</div>
			<button type="submit" class="w-full bg-white text-black font-black py-4 rounded-xl hover:scale-[1.02] transition-transform mt-6">Create Account</button>
			<p class="text-center text-zinc-500 text-sm mt-6 font-medium">Already have an account? <a href="/login" class="text-blue-400 hover:text-blue-300 font-bold">Log in</a></p>
		`
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>%s - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
	<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700;800;900&display=swap" rel="stylesheet">
	<style>
		body { font-family: 'Inter', sans-serif; }
	</style>
</head>
<body class="bg-[#09090b] text-white flex items-center justify-center min-h-screen p-6 bg-mesh">
	<div class="w-full max-w-md">
		<!-- Logo -->
		<div class="text-center mb-8">
			<a href="/" class="text-3xl font-black tracking-tighter">Nyumba.</a>
			<p class="text-zinc-500 text-sm mt-2 font-medium">Find your sanctuary</p>
		</div>
		
		<!-- Form Card -->
		<div class="bg-white/5 border border-white/10 rounded-[2rem] p-8 md:p-10 backdrop-blur-xl">
			<h2 class="text-3xl font-black tracking-tighter mb-2 text-center">%s</h2>
			<p class="text-zinc-500 text-center mb-8 text-sm font-medium">%s</p>
			
			<form action="/%s" method="POST" class="space-y-4">
				%s
			</form>
			
			<!-- Social Login Divider -->
			<div class="relative my-8">
				<div class="absolute inset-0 flex items-center">
					<div class="w-full border-t border-white/10"></div>
				</div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="bg-[#09090b] px-2 text-zinc-500 font-bold tracking-widest">Or continue with</span>
				</div>
			</div>
			
			<!-- Google Button -->
			<button type="button" class="w-full bg-white/5 border border-white/10 rounded-xl py-3 px-4 flex items-center justify-center gap-3 hover:bg-white/10 transition-colors font-bold text-sm">
				<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
					<path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
					<path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
					<path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
					<path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
				</svg>
				Google
			</button>
		</div>
		
		<!-- Footer -->
		<p class="text-center text-zinc-600 text-xs mt-8">
			By %s, you agree to our <a href="#" class="text-zinc-400 hover:text-white transition-colors">Terms</a> and <a href="#" class="text-zinc-400 hover:text-white transition-colors">Privacy Policy</a>
		</p>
	</div>
	
	<style>
		.bg-mesh {
			background-image: 
				radial-gradient(at 0%% 0%%, rgba(30, 58, 138, 0.15) 0px, transparent 50%%),
				radial-gradient(at 100%% 0%%, rgba(20, 184, 166, 0.1) 0px, transparent 50%%),
				radial-gradient(at 100%% 100%%, rgba(30, 58, 138, 0.1) 0px, transparent 50%%);
		}
	</style>
</body>
</html>`, mode, mode, 
		map[bool]string{true: "Welcome back to your sanctuary", false: "Join Kenya's premier housing platform"}[isLogin],
		map[bool]string{true: "login", false: "signup"}[isLogin],
		formFields,
		map[bool]string{true: "signing in", false: "signing up"}[isLogin])
}

func GetStaticHTML(title, content string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>%s - Nyumba</title>
	<script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-[#09090b] text-white">
	%s
	<main class="max-w-3xl mx-auto px-6 pt-48 pb-32">
		<h1 class="text-6xl md:text-8xl font-black tracking-tighter mb-12">%s</h1>
		<p class="text-zinc-400 font-medium leading-relaxed">%s</p>
	</main>
	%s
</body>
</html>`, title, getHeader(), title, content, getFooter())
}