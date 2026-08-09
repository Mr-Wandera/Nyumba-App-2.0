
import express from 'express';
import fs from 'fs';
import path from 'path';
import { KenyaLocationsData } from './locations';
import { 
    getLandingHTML, 
    getExploreHTML, 
    getLandlordHTML, 
    getAuthHTML, 
    getStaticHTML,
    getPropertyDetailHTML
} from './templates';

const app = express();
const PORT = 3000;

app.use(express.json());

// Helper to read houses
const getHouses = () => {
    try {
        const data = fs.readFileSync(path.join(process.cwd(), 'houses.json'), 'utf-8');
        return JSON.parse(data);
    } catch (err) {
        return [];
    }
};

// Helper to save houses
const saveHouses = (houses: any[]) => {
    fs.writeFileSync(path.join(process.cwd(), 'houses.json'), JSON.stringify(houses, null, 2));
};

// Helper to read users
const getUsers = () => {
    try {
        const data = fs.readFileSync(path.join(process.cwd(), 'users.json'), 'utf-8');
        return JSON.parse(data);
    } catch (err) {
        return [];
    }
};

// Helper to save users
const saveUsers = (users: any[]) => {
    fs.writeFileSync(path.join(process.cwd(), 'users.json'), JSON.stringify(users, null, 2));
};

// Helper to parse cookies
const parseCookies = (req: express.Request) => {
    const list: Record<string, string> = {};
    const rc = req.headers.cookie;
    if (rc) {
        rc.split(';').forEach(cookie => {
            const parts = cookie.split('=');
            if (parts.length >= 2) {
                list[parts[0].trim()] = decodeURIComponent(parts.slice(1).join('='));
            }
        });
    }
    return list;
};

// Pages
app.get('/', (req, res) => res.send(getLandingHTML()));
app.get('/explore', (req, res) => res.send(getExploreHTML()));
app.get('/landlord', (req, res) => res.send(getLandlordHTML()));
app.get('/login', (req, res) => res.send(getAuthHTML('login')));
app.get('/signup', (req, res) => res.send(getAuthHTML('signup')));

// Auth API Routes
app.post('/api/auth/login', (req, res) => {
    const { email, password } = req.body;
    if (!email || !password) {
        return res.status(400).json({ error: "Email and password are required" });
    }
    const users = getUsers();
    const cleanEmail = email.trim().toLowerCase();
    const user = users.find((u: any) => u.email.toLowerCase() === cleanEmail);
    if (!user) {
        return res.status(401).json({ error: "Invalid email or password" });
    }
    if (user.password && user.password !== password) {
        return res.status(401).json({ error: "Invalid email or password" });
    }

    const userData = {
        id: user.id,
        name: user.name || user.username || user.email.split('@')[0],
        email: user.email,
        phone: user.phone || "+254700000000",
        role: user.role || "tenant"
    };

    res.setHeader('Set-Cookie', `nyumba_token=${encodeURIComponent(JSON.stringify(userData))}; Path=/; HttpOnly; SameSite=Lax; Max-Age=86400`);
    return res.json({
        status: "success",
        user: userData
    });
});

app.post('/api/auth/register', (req, res) => {
    const { name, email, phone, password, role } = req.body;
    if (!email || !password || !name) {
        return res.status(400).json({ error: "Name, email, and password are required" });
    }
    const cleanEmail = email.trim().toLowerCase();
    const users = getUsers();
    if (users.some((u: any) => u.email.toLowerCase() === cleanEmail)) {
        return res.status(400).json({ error: "An account with this email address already exists" });
    }

    const newUser = {
        id: users.length + 1,
        name: name.trim(),
        username: name.trim().toLowerCase().replace(/\s+/g, ''),
        email: cleanEmail,
        phone: phone ? phone.trim() : "+254700000000",
        password: password,
        role: role || "tenant",
        created_at: new Date().toISOString()
    };

    users.push(newUser);
    saveUsers(users);

    const userData = {
        id: newUser.id,
        name: newUser.name,
        email: newUser.email,
        phone: newUser.phone,
        role: newUser.role
    };

    res.setHeader('Set-Cookie', `nyumba_token=${encodeURIComponent(JSON.stringify(userData))}; Path=/; HttpOnly; SameSite=Lax; Max-Age=86400`);
    return res.status(201).json({
        status: "success",
        user: userData
    });
});

app.get('/api/auth/me', (req, res) => {
    const cookies = parseCookies(req);
    const token = cookies.nyumba_token;
    if (token) {
        try {
            const user = JSON.parse(token);
            return res.json({ authenticated: true, user });
        } catch (e) {}
    }
    return res.status(401).json({ authenticated: false });
});

app.post('/api/auth/logout', (req, res) => {
    res.setHeader('Set-Cookie', `nyumba_token=; Path=/; HttpOnly; Max-Age=0`);
    return res.json({ status: "success", message: "Logged out successfully" });
});

app.get('/logout', (req, res) => {
    res.setHeader('Set-Cookie', `nyumba_token=; Path=/; HttpOnly; Max-Age=0`);
    return res.redirect('/login');
});

app.get('/property/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const houses = getHouses();
    const house = houses.find((h: any) => h.id === id);
    if (!house) {
        return res.redirect('/explore');
    }
    res.send(getPropertyDetailHTML(house));
});

app.get('/about', (req, res) => res.send(getStaticHTML('About Us', `
    <p class="mb-4">Nyumba is Kenya's premier house-hunting platform, dedicated to making the process of finding a home transparent, secure, and efficient.</p>
    <p class="mb-4">We believe that everyone deserves a home they love, without the hassle of unreliable agents and hidden fees. By connecting tenants directly with verified landlords, we ensure a smoother transition for everyone involved.</p>
    <p>Our team works tirelessly to verify every listing, providing you with peace of mind as you search for your next sanctuary.</p>
`)));

app.get('/contact', (req, res) => res.send(getStaticHTML('Contact Us', `
    <p class="mb-6">Have questions or need assistance? We're here to help!</p>
    <div class="space-y-4">
        <div>
            <h3 class="font-bold text-gray-900">Email</h3>
            <p>support@nyumba.co.ke</p>
        </div>
        <div>
            <h3 class="font-bold text-gray-900">Phone</h3>
            <p>+254 700 000 000</p>
        </div>
        <div>
            <h3 class="font-bold text-gray-900">Office</h3>
            <p>Nairobi, Kenya</p>
        </div>
    </div>
`)));

// Location API Routes
app.get('/api/locations/counties', (req, res) => {
    const counties = KenyaLocationsData.map(c => ({
        id: c.code,
        code: c.code,
        name: c.name
    }));
    res.json(counties);
});

app.get('/api/locations/sub-counties', (req, res) => {
    const { county_id } = req.query;
    if (!county_id) {
        const allSubs: any[] = [];
        KenyaLocationsData.forEach(c => {
            c.subCounties.forEach((s, sIdx) => {
                allSubs.push({
                    id: c.code * 100 + sIdx + 1,
                    county_id: c.code,
                    name: s.name
                });
            });
        });
        return res.json(allSubs);
    }
    const countyCode = parseInt(county_id as string);
    const county = KenyaLocationsData.find(c => c.code === countyCode);
    if (!county) return res.json([]);
    const subCounties = county.subCounties.map((s, idx) => ({
        id: county.code * 100 + idx + 1,
        county_id: county.code,
        name: s.name
    }));
    res.json(subCounties);
});

app.get('/api/locations/towns', (req, res) => {
    const { county_id } = req.query;
    if (!county_id) {
        const allTowns: any[] = [];
        KenyaLocationsData.forEach(c => {
            c.towns.forEach((t, tIdx) => {
                allTowns.push({
                    id: c.code * 100 + tIdx + 1,
                    county_id: c.code,
                    name: t
                });
            });
        });
        return res.json(allTowns);
    }
    const countyCode = parseInt(county_id as string);
    const county = KenyaLocationsData.find(c => c.code === countyCode);
    if (!county) return res.json([]);
    const towns = county.towns.map((t, idx) => ({
        id: county.code * 100 + idx + 1,
        county_id: county.code,
        name: t
    }));
    res.json(towns);
});

app.get('/api/locations/wards', (req, res) => {
    const { sub_county_id } = req.query;
    if (!sub_county_id) return res.json([]);
    const subId = parseInt(sub_county_id as string);
    const countyCode = Math.floor(subId / 100);
    const subIdx = (subId % 100) - 1;
    const county = KenyaLocationsData.find(c => c.code === countyCode);
    if (!county || subIdx < 0 || subIdx >= county.subCounties.length) return res.json([]);
    const sub = county.subCounties[subIdx];
    const wards = sub.wards.map((w, idx) => ({
        id: subId * 100 + idx + 1,
        sub_county_id: subId,
        name: w
    }));
    res.json(wards);
});

app.get('/api/locations/neighborhoods', (req, res) => {
    const { sub_county_id } = req.query;
    if (!sub_county_id) return res.json([]);
    const subId = parseInt(sub_county_id as string);
    const countyCode = Math.floor(subId / 100);
    const subIdx = (subId % 100) - 1;
    const county = KenyaLocationsData.find(c => c.code === countyCode);
    if (!county || subIdx < 0 || subIdx >= county.subCounties.length) return res.json([]);
    const sub = county.subCounties[subIdx];
    const neighborhoods = sub.neighborhoods.map((n, idx) => ({
        id: subId * 1000 + idx + 1,
        sub_county_id: subId,
        name: n
    }));
    res.json(neighborhoods);
});

app.get('/api/locations/hierarchy', (req, res) => {
    res.json(KenyaLocationsData);
});

// House API Routes
app.get('/api/houses', (req, res) => {
    const { search, maxPrice, county_id, sub_county_id, town_id } = req.query;
    let houses = getHouses();

    if (county_id) {
        const cId = parseInt(county_id as string);
        const countyObj = KenyaLocationsData.find(c => c.code === cId);
        if (countyObj) {
            const countyName = countyObj.name.toLowerCase();
            houses = houses.filter((h: any) => 
                h.county_id === cId || 
                (h.location && h.location.toLowerCase().includes(countyName))
            );
        }
    }

    if (sub_county_id) {
        const sId = parseInt(sub_county_id as string);
        const countyCode = Math.floor(sId / 100);
        const subIdx = (sId % 100) - 1;
        const countyObj = KenyaLocationsData.find(c => c.code === countyCode);
        if (countyObj && subIdx >= 0 && subIdx < countyObj.subCounties.length) {
            const subName = countyObj.subCounties[subIdx].name.toLowerCase();
            houses = houses.filter((h: any) => 
                h.sub_county_id === sId || 
                (h.location && h.location.toLowerCase().includes(subName))
            );
        }
    }

    if (town_id) {
        const tId = parseInt(town_id as string);
        const countyCode = Math.floor(tId / 100);
        const townIdx = (tId % 100) - 1;
        const countyObj = KenyaLocationsData.find(c => c.code === countyCode);
        if (countyObj && townIdx >= 0 && townIdx < countyObj.towns.length) {
            const townName = countyObj.towns[townIdx].toLowerCase();
            houses = houses.filter((h: any) => 
                h.town_id === tId || 
                (h.location && h.location.toLowerCase().includes(townName))
            );
        }
    }

    if (search) {
        const s = (search as string).toLowerCase();
        houses = houses.filter((h: any) => 
            (h.building_name && h.building_name.toLowerCase().includes(s)) || 
            (h.location && h.location.toLowerCase().includes(s))
        );
    }

    if (maxPrice) {
        const price = parseInt(maxPrice as string);
        houses = houses.filter((h: any) => h.price <= price);
    }

    res.json(houses);
});

app.post('/api/add-house', (req, res) => {
    const houses = getHouses();

    // Construct nice location string if county/sub-county provided
    let fullLocation = req.body.location || '';
    if (req.body.county_id) {
        const cObj = KenyaLocationsData.find(c => c.code === parseInt(req.body.county_id));
        if (cObj) {
            const locParts = [];
            if (req.body.location) locParts.push(req.body.location);
            if (req.body.sub_county_id) {
                const sIdx = (parseInt(req.body.sub_county_id) % 100) - 1;
                if (cObj.subCounties[sIdx]) locParts.push(cObj.subCounties[sIdx].name);
            }
            locParts.push(cObj.name);
            fullLocation = locParts.join(', ');
        }
    }

    const newHouse = {
        id: houses.length + 1,
        ...req.body,
        location: fullLocation,
        county_id: req.body.county_id ? parseInt(req.body.county_id) : undefined,
        sub_county_id: req.body.sub_county_id ? parseInt(req.body.sub_county_id) : undefined,
        town_id: req.body.town_id ? parseInt(req.body.town_id) : undefined,
        price: req.body.price ? parseInt(req.body.price) : 0,
        is_paid: false,
        image_urls: [`https://picsum.photos/seed/${Math.random()}/800/600`],
        bedrooms: req.body.bedrooms ? parseInt(req.body.bedrooms) : Math.floor(Math.random() * 3) + 1,
        bathrooms: req.body.bathrooms ? parseInt(req.body.bathrooms) : Math.floor(Math.random() * 2) + 1,
        landlord_phone: "+254700000000",
        description: req.body.description || "A newly listed property on Nyumba."
    };
    houses.push(newHouse);
    saveHouses(houses);
    res.status(201).json({ message: 'Property listed successfully!', house: newHouse });
});

app.post('/api/trigger-payment', (req, res) => {
    const { houseId } = req.body;
    // Initiate M-Pesa STK trigger
    setTimeout(() => {
        res.json({ 
            success: true, 
            message: `M-Pesa STK Push sent for House #${houseId}. Please check your phone to complete the KES 1,000 payment.` 
        });
    }, 1000);
});

app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server running on http://localhost:${PORT}`);
});

